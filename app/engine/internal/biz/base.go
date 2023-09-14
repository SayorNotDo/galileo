package biz

import "fmt"

const (
	TypeDefaultJob  = "job:default"
	TypeDelayedJob  = "job:delayed"
	TypePeriodicJob = "job:periodic"
)

func QueuePrefix(qName string) string {
	return fmt.Sprintf("asynq:{%s}:", qName)
}

func ScheduledKey(qName string) string {
	return fmt.Sprintf("%sscheduled", QueuePrefix(qName))
}

func JobKeyPrefix(qName string) string {
	return fmt.Sprintf("%st:", QueuePrefix(qName))
}

func JobKey(qName, jobID string) string {
	return fmt.Sprintf("%s%s", JobKeyPrefix(qName), jobID)
}
