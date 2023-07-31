package constant

import "github.com/golang-jwt/jwt/v4"

var (
	KafkaBrokers = []string{"localhost:9092"}
	KafkaTopic   = "test_data_original_report"
	FieldList    = []string{"#report_type", "#sync", "#timestamp", "#execute_id", "properties"}
	ReportType   = []string{"track", "monitor"}
	ReportKey    = []byte("test_data_original_report_secret_key")
)

type ReportClaims struct {
	Machine string `json:"machine,omitempty"`
	jwt.RegisteredClaims
}
