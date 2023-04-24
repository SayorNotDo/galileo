package data

import (
	"bytes"
	"context"
	"galileo/pkg/errResponse"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/errors"
	"strings"

	"galileo/app/file/internal/biz"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/go-kratos/kratos/v2/log"
)

var OssSessionName = "kratos-base-project"

type fileRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileRepo .
func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "dataService.fileRepo")),
	}
}

func (r *fileRepo) UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	client, err := oss.New(r.data.config.Oss.Endpoint, r.data.config.Oss.AccessKey, r.data.config.Oss.AccessSecret)
	if err != nil {
		return "", errResponse.SetErrByReason(errResponse.ReasonOssConfigWrong)
	}

	bucket, err := client.Bucket(r.data.config.Oss.BucketName)
	if err != nil {
		return "", errResponse.SetErrByReason(errResponse.ReasonOssConfigWrong)
	}
	path := "uploadFile/" + strings.Trim(fileType, ".") + "/" + fileName + fileType
	err = bucket.PutObject(path, bytes.NewReader(content))
	if err != nil {
		return "", errResponse.SetErrByReason(errResponse.ReasonOssPutObjectFail)
	}
	url := "https://" + r.data.config.Oss.BucketName + "." + r.data.config.Oss.Endpoint + "/" + path
	return url, nil
}

func (r *fileRepo) GetOssStsToken(ctx context.Context) (*biz.OssStsToken, error) {
	config := &openapi.Config{
		AccessKeyId:     &r.data.config.Oss.AccessKey,
		AccessKeySecret: &r.data.config.Oss.AccessSecret,
	}

	config.Endpoint = tea.String("file-server-domain")
	client, err := sts.NewClient(config)
	if err != nil {
		return &biz.OssStsToken{}, errors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	assumeRoleRequest := &sts.AssumeRoleRequest{
		RoleArn:         &r.data.config.Oss.StsRoleArn,
		RoleSessionName: &OssSessionName,
	}
	res, err := client.AssumeRole(assumeRoleRequest)
	if err != nil {
		return &biz.OssStsToken{}, errors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	return &biz.OssStsToken{
		AccessKey:     *res.Body.Credentials.AccessKeyId,
		AccessSecret:  *res.Body.Credentials.AccessKeySecret,
		Expiration:    *res.Body.Credentials.Expiration,
		SecurityToken: *res.Body.Credentials.SecurityToken,
		EndPoint:      r.data.config.Oss.Endpoint,
		BucketName:    r.data.config.Oss.BucketName,
		Region:        r.data.config.Oss.Region,
		Url:           "https://" + r.data.config.Oss.BucketName + "." + r.data.config.Oss.Endpoint + "/",
	}, nil
}
