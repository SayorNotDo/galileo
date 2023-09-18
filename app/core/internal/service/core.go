package service

import (
	"context"
	"encoding/json"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/pkg/constant"
	. "galileo/app/core/internal/pkg/middleware/auth"
	"galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *CoreService) ExecuteToken(ctx context.Context, req *v1.ExecuteTokenRequest) (*v1.ExecuteTokenReply, error) {
	token, err := c.cc.ExecuteToken(ctx, req.Machine)
	if err != nil {
		return nil, err
	}
	return &v1.ExecuteTokenReply{
		ExecuteToken: token,
	}, nil
}

/*TrackReportData 数据上报后台接口 */
func (c *CoreService) TrackReportData(ctx context.Context, req *v1.TrackReportDataRequest) (*emptypb.Empty, error) {
	/* 数据上报接口单独鉴权逻辑 */
	var claims *ReportClaims
	if tr, ok := transport.FromServerContext(ctx); ok {
		if ht, ok := tr.(*http.Transport); ok {
			/* 实际生产环境中，需要根据具体情况考虑是否需要处理“X-Forwarded-For”头部以获取真实的客户端IP地址 */
			ctx = context.WithValue(ctx, "RemoteAddr", ht.Request().RemoteAddr)
		}
		auth := tr.RequestHeader().Get("x-execute-token")
		/* 解析x-execute-token 获取相关信息进行校验 */
		token, err := jwt.ParseWithClaims(auth, &ReportClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.ReportKey), nil
		})
		if err != nil {
			return nil, err
		}
		claims, ok = token.Claims.(*ReportClaims)
		if !ok || !token.Valid {
			return nil, errResponse.SetErrByReason(errResponse.ReasonUserUnauthorized)
		}
	}
	originData := req.Data
	var dataList []map[string]interface{}
	/* 解析接口上报的数据 */
	if err := json.Unmarshal([]byte(originData), &dataList); err != nil {
		return nil, err
	}
	/* 调用业务函数 */
	if err := c.cc.DataReportTrack(ctx, dataList, claims); err != nil {
		return nil, err
	}
	return nil, nil
}
