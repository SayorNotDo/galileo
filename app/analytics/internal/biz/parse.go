package biz

import (
	v1 "galileo/api/analytics/v1"
)

func ParseEvent(event *v1.EventAnalyticsRequest_Event) string {
	switch event.Analysis {
	case "A100":
		return `COUNT(` + event.EventUUID + `)`
	case "A101":
	default:
		return ""
	}
	return ""
}
