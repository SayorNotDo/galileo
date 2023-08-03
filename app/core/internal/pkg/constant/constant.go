package constant

import "time"

var (
	KafkaBrokers = []string{"localhost:9092"}
	KafkaTopic   = "test_data_original_report"
	FieldList    = []string{"#report_type", "#sync", "#timestamp", "#execute_id", "properties"}
	ReportType   = []string{"track", "monitor"}
	ReportKey    = "test_data_original_report_secret_key"
)

const (
	RedisMaxRetries = 10
	TaskProgressKey = "taskProgress"
	TaskExpiration  = time.Hour * 24 * 30
)
