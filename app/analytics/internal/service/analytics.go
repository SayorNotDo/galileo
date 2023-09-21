package service

import (
	"context"
	v1 "galileo/api/analytics/v1"
)

func (s *AnalyticsService) EventAnalytics(ctx context.Context, req *v1.EventAnalyticsRequest) (*v1.EventAnalyticReply, error) {
	/* 解析指标 */
	event := req.Events
	s.log.Debugf("EventAnalytics: %s ,%s, %s, %d", event.Analysis, event.EventUUID, event.Relation, event.Type)
	/* 基于UUID获取元事件列表中的相关信息 */
	/* 参数校验 */
	/* SQL生成 */
	return nil, nil
}
