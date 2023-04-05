package ctxdata

import "context"

const UserIdKey = "kratos-userId"
const Username = "kratos-username"

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
