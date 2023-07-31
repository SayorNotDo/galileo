package constant

var (
	KafkaBrokers = []string{"localhost:9092"}
	KafkaTopic   = "test_data_original_report"
	FieldList    = []string{"#report_type", "#sync", "#timestamp", "#execute_id", "properties"}
	ReportType   = []string{"track", "monitor"}
)
