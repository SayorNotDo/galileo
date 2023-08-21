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
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *CoreService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	// add trace
	tr := otel.Tracer("scheduler")
	ctx, span := tr.Start(ctx, "get user info")
	span.SpanContext()
	defer span.End()
	return c.cc.CreateUser(ctx, req)
}

func (c *CoreService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	// add trace
	tr := otel.Tracer("scheduler")
	ctx, span := tr.Start(ctx, "login")
	span.SpanContext()
	defer span.End()
	return c.cc.Login(ctx, req)
}

func (c *CoreService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	// add trace
	tr := otel.Tracer("scheduler")
	ctx, span := tr.Start(ctx, "logout")
	span.SpanContext()
	defer span.End()
	return c.cc.Logout(ctx)
}

func (c *CoreService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	return c.cc.UpdatePassword(ctx, req)
}

func (c *CoreService) UserDetail(ctx context.Context, req *emptypb.Empty) (*v1.UserDetailReply, error) {
	return c.cc.UserDetail(ctx, &emptypb.Empty{})
}

func (c *CoreService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return c.cc.ListUser(ctx, req.PageNum, req.PageSize)
}

func (c *CoreService) UpdateUserInfo(ctx context.Context, req *v1.UserInfoUpdateRequest) (*v1.UserInfoUpdateReply, error) {
	return c.cc.UpdateUserInfo(ctx, req)
}

func (c *CoreService) DeleteUser(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteReply, error) {
	return c.cc.DeleteUser(ctx, req.Id)
}

func (c *CoreService) ExecuteToken(ctx context.Context, req *v1.ExecuteTokenRequest) (*v1.ExecuteTokenReply, error) {
	token, err := c.cc.ExecuteToken(ctx, req.Machine)
	if err != nil {
		return nil, err
	}
	return &v1.ExecuteTokenReply{
		ExecuteToken: token,
	}, nil
}

/*DataReportTrack 数据上报后台接口 */
func (c *CoreService) DataReportTrack(ctx context.Context, req *v1.DataReportTrackRequest) (*emptypb.Empty, error) {
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
