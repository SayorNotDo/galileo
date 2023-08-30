package ctxdata

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strconv"
)

const (
	UserIdKey         = "x-md-global-userId"
	Username          = "x-md-global-username"
	AuthorizationKey  = "Authorization"
	AuthorizationType = "Bearer"
	Issuer            = "Galileo"
)

func GetUserId(ctx context.Context) uint32 {
	if v, ok := ctx.Value(UserIdKey).(uint32); ok {
		return v
	}
	return 0
}

func GetUserName(ctx context.Context) string {
	if v, ok := ctx.Value(Username).(string); ok {
		return v
	}
	return ""
}

func MethodFromContext(ctx context.Context) string {
	if tr, ok := transport.FromServerContext(ctx); ok {
		return tr.(*http.Transport).Request().Method
	}
	return ""
}

func UserIdFromMetaData(ctx context.Context) uint32 {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0
	}
	uidStr := md.Get(UserIdKey)
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		return 0
	}
	return uint32(uid)
}
