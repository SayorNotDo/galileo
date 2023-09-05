package biz

type DelayedJobPayload struct {
	Task    int64  `json:"task"`
	Worker  int64  `json:"worker"`
	Message string `json:"message"`
}
