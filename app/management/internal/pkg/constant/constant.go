package constant

import "time"

const (
	RedisMaxRetries = 10
	TaskProgressKey = "taskProgress"
	TaskExpiration  = time.Hour * 24 * 30
)

const (
	DEFAULT = iota
	DELAYED
	PERIODIC
)
