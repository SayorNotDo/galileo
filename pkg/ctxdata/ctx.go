package ctxdata

import "context"

const (
	UserIdKey         = "x-md-global-userId"
	Username          = "x-md-global-username"
	AuthorizationKey  = "Authorization"
	AuthorizationType = "Bearer"
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
