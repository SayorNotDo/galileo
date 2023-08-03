package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	v1 "galileo/api/core/v1"
	userService "galileo/api/user/v1"
	"galileo/app/core/internal/biz"
	. "galileo/app/core/internal/pkg/constant"
	. "galileo/app/core/internal/pkg/middleware/auth"
	"galileo/pkg/encryption"
	. "galileo/pkg/errResponse"
	. "galileo/pkg/factory"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

type coreRepo struct {
	data *Data
	log  *log.Helper
}

func NewCoreRepo(data *Data, logger log.Logger) biz.CoreRepo {
	return &coreRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "core.Repo")),
	}
}

func (r *coreRepo) UserByUsername(c context.Context, username string) (*biz.User, error) {
	user, err := r.data.uc.GetUserByUsername(c, &userService.UsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       user.Id,
		Phone:    user.Phone,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

func (r *coreRepo) UserById(c context.Context, id uint32) (*biz.User, error) {
	user, err := r.data.uc.GetUser(c, &userService.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Nickname: user.Nickname,
		Email:    user.Email,
	}, nil
}

func (r *coreRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	createUser, err := r.data.uc.CreateUser(c, &userService.CreateUserRequest{
		Username: user.Nickname,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:        createUser.Id,
		Username:  createUser.Username,
		CreatedAt: createUser.CreatedAt.AsTime(),
	}, nil
}

func (r *coreRepo) UpdatePassword(ctx context.Context, password string) (bool, error) {
	if _, err := r.data.uc.UpdatePassword(ctx, &userService.UpdatePasswordRequest{NewPassword: password}); err != nil {
		return false, err
	}
	return true, nil
}

func (r *coreRepo) VerifyPassword(c context.Context, password, encryptedPassword string) (bool, error) {
	if ret, err := r.data.uc.VerifyPassword(c, &userService.VerifyPasswordRequest{Password: password, HashedPassword: encryptedPassword}); err != nil {
		return false, err
	} else {
		return ret.Success, nil
	}
}

func (r *coreRepo) ListUser(c context.Context, pageNum, pageSize int32) ([]*v1.UserDetail, int32, error) {
	rsp, err := r.data.uc.ListUser(c, &userService.ListUserRequest{PageNum: pageNum, PageSize: pageSize})
	if err != nil {
		return nil, 0, err
	}
	rv := make([]*v1.UserDetail, 0)
	for _, u := range rsp.Data {
		rv = append(rv, &v1.UserDetail{
			Id:          u.Id,
			Nickname:    u.Nickname,
			ChineseName: u.ChineseName,
			Phone:       u.Phone,
			Email:       u.Email,
			Role:        u.Role,
		})
	}
	total := rsp.Total
	return rv, total, nil
}

func (r *coreRepo) SoftDeleteUser(c context.Context, uid uint32) (bool, error) {
	rsp, err := r.data.uc.SoftDeleteUser(c, &userService.SoftDeleteRequest{Id: uid})
	if err != nil {
		return false, err
	}
	return rsp.Deleted, nil
}

func (r *coreRepo) UpdateUserInfo(c context.Context, user *biz.User) (bool, error) {
	return false, nil
}

func (r *coreRepo) SetToken(ctx context.Context, token string) (string, error) {
	key := encryption.EncodeMD5(token)
	r.data.redisCli.Set(ctx, "token:"+key, token, TokenExpiration)
	return key, nil
}

func (r *coreRepo) DestroyToken(ctx context.Context) error {
	tr, _ := transport.FromServerContext(ctx)
	jwtToken := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)[1]
	key := encryption.EncodeMD5(jwtToken)
	r.data.redisCli.Del(ctx, "token:"+key)
	return nil
}

func (r *coreRepo) DataReportTrack(ctx context.Context, data []map[string]interface{}, claims *ReportClaims) error {
	/* 数据清洗 */
	var cleanDataList []map[string]interface{}
	for _, v := range data {
		cleanData, err := dataCleaner(ctx, v, claims)
		if err != nil {
			return err
		}
		cleanDataList = append(cleanDataList, cleanData)
	}
	/* 写入Kafka队列 */
	if err := sendToKafka(r.data.kafkaProducer, KafkaTopic, cleanDataList); err != nil {
		return err
	}
	/* 写入Redis缓存 */
	if err := r.sendToRedis(ctx, cleanDataList); err != nil {
		/* 写入失败时增加重试/缓存机制，保证与Kafka队列消息相同 */
		return err
	}
	return nil
}

func dataCleaner(ctx context.Context, data map[string]interface{}, claim *ReportClaims) (map[string]interface{}, error) {
	/* 检查必需字段是否存在缺失情况 */
	for _, v := range FieldList {
		value, ok := data[v]
		if !ok {
			return nil, errors.New("missing required field: " + v)
		}
		/* 数据预处理噪音、错误、缺失值和异常值 */
		switch v {
		/* 时间戳格式是否合法 */
		case "#timestamp":
			tNum, err := strconv.ParseInt(value.(string), 10, 64)
			if err != nil {
				return nil, errors.New("#timestamp's value convert from interface{} to int64 failed")
			}
			if t := time.Unix(tNum, 0); t.IsZero() {
				return nil, errors.New("#timestamp's value illegal")
			}
			/* 是否存在时间穿越 */
			currentTime := time.Now().UnixNano() / 1e6
			if tNum > currentTime {
				return nil, errors.New("#timestamp's value cannot be greater than current time")
			}
		/* 上报类型是否合法 */
		case "#report_type":
			reportType, ok := value.(string)
			if !ok {
				return nil, errors.New("#report_type' value convert from interface{} to string failed")
			}
			for _, v := range ReportType {
				if reportType == v {
					break
				}
				return nil, errors.New("#report_type' value illegal")
			}
		/* 其余字段校验是否为空 */
		case "#sync", "#execute_id":
			field, ok := value.(string)
			if !ok {
				return nil, errors.New(v + "'s value convert from interface{} to string failed")
			}
			if field == "" {
				return nil, errors.New(v + "'s value cannot be empty")
			}
		}
	}
	/* 服务端属性添加 */
	remoteAddr := ctx.Value("RemoteAddr").(string)
	fieldToAdd := map[string]interface{}{
		"$machine":        claim.Machine,
		"$remote_address": remoteAddr,
	}
	properties := data["properties"].(map[string]interface{})
	for k, v := range fieldToAdd {
		properties[k] = v
	}
	/* 数据去重 */
	/* 数据转换 */
	/* 数据过滤 */
	return data, nil
}

func sendToKafka(producer sarama.SyncProducer, topic string, data []map[string]interface{}) error {
	/* 数据转换为JSON字节 */
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.New("Failed to marshal data to JSON: " + err.Error())
	}
	/* 发送消息至Kafka */
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jsonData),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return errors.New("Failed to send message to Kafka: " + err.Error())
	}
	fmt.Printf("Message sent to partition %d, offset %d\n", partition, offset)
	return nil
}

func (r *coreRepo) sendToRedis(ctx context.Context, data []map[string]interface{}) error {
	for _, item := range data {
		switch item["#report_type"] {
		case "track":
			properties, ok := item["properties"].(map[string]interface{})
			if !ok {
				return errors.New("invalid properties structure")
			}
			taskName, ok := properties["#task_name"]
			if !ok {
				return SetCustomizeErrMsg(ReasonParamsError, "properties do not define #task_name attribute")
			}
			executeId, err := strconv.ParseInt(item["#execute_id"].(string), 10, 64)
			if err != nil {
				return SetCustomizeErrMsg(ReasonParamsError, "convert #execute_id attribute to int64 error: "+err.Error())
			}
			/* 构建Redis Key */
			taskKey := NewTaskProgressKey(taskName.(string), executeId)
			/* map[string]interface{} To string */
			jsonItem, err := json.Marshal(item)
			if err != nil {
				return SetCustomizeErrMsg(ReasonSystemError, "json marshal error: "+err.Error())
			}
			if err := r.RedisLPushTask(ctx, taskKey, string(jsonItem)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *coreRepo) RedisLPushTask(ctx context.Context, key, val string) error {
	/* 事务函数: 通过乐观锁实现任务数据写入Redis */
	txf := func(tx *redis.Tx) error {
		err := r.data.redisCli.LPush(ctx, key, val).Err()
		if err != nil {
			return err
		}
		return nil
	}
	for i := 0; i < RedisMaxRetries; i++ {
		err := r.data.redisCli.Watch(ctx, txf, key)
		switch err {
		case nil:
			return nil
		case redis.TxFailedErr:
			continue
		}
	}
	return SetCustomizeErrMsg(ReasonSystemError, "max retries exceeded")
}
