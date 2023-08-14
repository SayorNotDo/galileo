// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/analytics/v1/analytics.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventAnalyticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events     []*EventAnalyticsRequest_Event    `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	EventsView *EventAnalyticsRequest_EventsView `protobuf:"bytes,2,opt,name=eventsView,proto3" json:"eventsView,omitempty"`
	Origin     string                            `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *EventAnalyticsRequest) Reset() {
	*x = EventAnalyticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_analytics_v1_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventAnalyticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventAnalyticsRequest) ProtoMessage() {}

func (x *EventAnalyticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_analytics_v1_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventAnalyticsRequest.ProtoReflect.Descriptor instead.
func (*EventAnalyticsRequest) Descriptor() ([]byte, []int) {
	return file_api_analytics_v1_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *EventAnalyticsRequest) GetEvents() []*EventAnalyticsRequest_Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *EventAnalyticsRequest) GetEventsView() *EventAnalyticsRequest_EventsView {
	if x != nil {
		return x.EventsView
	}
	return nil
}

func (x *EventAnalyticsRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

type EventAnalyticReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAxis []string `protobuf:"bytes,1,rep,name=xAxis,proto3" json:"xAxis,omitempty"`
	YAxis []string `protobuf:"bytes,2,rep,name=yAxis,proto3" json:"yAxis,omitempty"`
}

func (x *EventAnalyticReply) Reset() {
	*x = EventAnalyticReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_analytics_v1_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventAnalyticReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventAnalyticReply) ProtoMessage() {}

func (x *EventAnalyticReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_analytics_v1_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventAnalyticReply.ProtoReflect.Descriptor instead.
func (*EventAnalyticReply) Descriptor() ([]byte, []int) {
	return file_api_analytics_v1_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *EventAnalyticReply) GetXAxis() []string {
	if x != nil {
		return x.XAxis
	}
	return nil
}

func (x *EventAnalyticReply) GetYAxis() []string {
	if x != nil {
		return x.YAxis
	}
	return nil
}

type EventAnalyticsRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventAnalyticsRequest_Filter) Reset() {
	*x = EventAnalyticsRequest_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_analytics_v1_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventAnalyticsRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventAnalyticsRequest_Filter) ProtoMessage() {}

func (x *EventAnalyticsRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_analytics_v1_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventAnalyticsRequest_Filter.ProtoReflect.Descriptor instead.
func (*EventAnalyticsRequest_Filter) Descriptor() ([]byte, []int) {
	return file_api_analytics_v1_analytics_proto_rawDescGZIP(), []int{0, 0}
}

type EventAnalyticsRequest_Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Analysis         string                          `protobuf:"bytes,1,opt,name=analysis,proto3" json:"analysis,omitempty"`
	AnalysisDesc     string                          `protobuf:"bytes,2,opt,name=analysisDesc,proto3" json:"analysisDesc,omitempty"`
	EventId          string                          `protobuf:"bytes,3,opt,name=eventId,proto3" json:"eventId,omitempty"`
	EventName        string                          `protobuf:"bytes,4,opt,name=eventName,proto3" json:"eventName,omitempty"`
	EventNameDisplay string                          `protobuf:"bytes,5,opt,name=eventNameDisplay,proto3" json:"eventNameDisplay,omitempty"`
	EventType        string                          `protobuf:"bytes,6,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Relation         int32                           `protobuf:"varint,7,opt,name=relation,proto3" json:"relation,omitempty"`
	Type             int32                           `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	Filters          []*EventAnalyticsRequest_Filter `protobuf:"bytes,9,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *EventAnalyticsRequest_Event) Reset() {
	*x = EventAnalyticsRequest_Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_analytics_v1_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventAnalyticsRequest_Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventAnalyticsRequest_Event) ProtoMessage() {}

func (x *EventAnalyticsRequest_Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_analytics_v1_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventAnalyticsRequest_Event.ProtoReflect.Descriptor instead.
func (*EventAnalyticsRequest_Event) Descriptor() ([]byte, []int) {
	return file_api_analytics_v1_analytics_proto_rawDescGZIP(), []int{0, 1}
}

func (x *EventAnalyticsRequest_Event) GetAnalysis() string {
	if x != nil {
		return x.Analysis
	}
	return ""
}

func (x *EventAnalyticsRequest_Event) GetAnalysisDesc() string {
	if x != nil {
		return x.AnalysisDesc
	}
	return ""
}

func (x *EventAnalyticsRequest_Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventAnalyticsRequest_Event) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *EventAnalyticsRequest_Event) GetEventNameDisplay() string {
	if x != nil {
		return x.EventNameDisplay
	}
	return ""
}

func (x *EventAnalyticsRequest_Event) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *EventAnalyticsRequest_Event) GetRelation() int32 {
	if x != nil {
		return x.Relation
	}
	return 0
}

func (x *EventAnalyticsRequest_Event) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *EventAnalyticsRequest_Event) GetFilters() []*EventAnalyticsRequest_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type EventAnalyticsRequest_EventsView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupBy          []string                        `protobuf:"bytes,1,rep,name=groupBy,proto3" json:"groupBy,omitempty"`
	TimeParticleSize string                          `protobuf:"bytes,2,opt,name=timeParticleSize,proto3" json:"timeParticleSize,omitempty"`
	TableSorts       []string                        `protobuf:"bytes,3,rep,name=tableSorts,proto3" json:"tableSorts,omitempty"`
	ByType           string                          `protobuf:"bytes,4,opt,name=byType,proto3" json:"byType,omitempty"`
	Filters          []*EventAnalyticsRequest_Filter `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty"`
	RecentDay        string                          `protobuf:"bytes,6,opt,name=recentDay,proto3" json:"recentDay,omitempty"`
}

func (x *EventAnalyticsRequest_EventsView) Reset() {
	*x = EventAnalyticsRequest_EventsView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_analytics_v1_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventAnalyticsRequest_EventsView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventAnalyticsRequest_EventsView) ProtoMessage() {}

func (x *EventAnalyticsRequest_EventsView) ProtoReflect() protoreflect.Message {
	mi := &file_api_analytics_v1_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventAnalyticsRequest_EventsView.ProtoReflect.Descriptor instead.
func (*EventAnalyticsRequest_EventsView) Descriptor() ([]byte, []int) {
	return file_api_analytics_v1_analytics_proto_rawDescGZIP(), []int{0, 2}
}

func (x *EventAnalyticsRequest_EventsView) GetGroupBy() []string {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *EventAnalyticsRequest_EventsView) GetTimeParticleSize() string {
	if x != nil {
		return x.TimeParticleSize
	}
	return ""
}

func (x *EventAnalyticsRequest_EventsView) GetTableSorts() []string {
	if x != nil {
		return x.TableSorts
	}
	return nil
}

func (x *EventAnalyticsRequest_EventsView) GetByType() string {
	if x != nil {
		return x.ByType
	}
	return ""
}

func (x *EventAnalyticsRequest_EventsView) GetFilters() []*EventAnalyticsRequest_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *EventAnalyticsRequest_EventsView) GetRecentDay() string {
	if x != nil {
		return x.RecentDay
	}
	return ""
}

var File_api_analytics_v1_analytics_proto protoreflect.FileDescriptor

var file_api_analytics_v1_analytics_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x06, 0x0a, 0x15, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x69, 0x65,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x56, 0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x1a,
	0x08, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0xc3, 0x02, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x44, 0x65, 0x73, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a,
	0xf2, 0x01, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x69, 0x65, 0x77, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x79, 0x22, 0x40, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x78, 0x41,
	0x78, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x78, 0x41, 0x78, 0x69, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x79, 0x41, 0x78, 0x69, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x79, 0x41, 0x78, 0x69, 0x73, 0x32, 0x90, 0x01, 0x0a, 0x09, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01,
	0x2a, 0x22, 0x16, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x31, 0x0a, 0x10, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x1b, 0x67, 0x61, 0x6c, 0x69, 0x6c, 0x65, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_analytics_v1_analytics_proto_rawDescOnce sync.Once
	file_api_analytics_v1_analytics_proto_rawDescData = file_api_analytics_v1_analytics_proto_rawDesc
)

func file_api_analytics_v1_analytics_proto_rawDescGZIP() []byte {
	file_api_analytics_v1_analytics_proto_rawDescOnce.Do(func() {
		file_api_analytics_v1_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_analytics_v1_analytics_proto_rawDescData)
	})
	return file_api_analytics_v1_analytics_proto_rawDescData
}

var file_api_analytics_v1_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_analytics_v1_analytics_proto_goTypes = []interface{}{
	(*EventAnalyticsRequest)(nil),            // 0: api.analytics.v1.EventAnalyticsRequest
	(*EventAnalyticReply)(nil),               // 1: api.analytics.v1.EventAnalyticReply
	(*EventAnalyticsRequest_Filter)(nil),     // 2: api.analytics.v1.EventAnalyticsRequest.Filter
	(*EventAnalyticsRequest_Event)(nil),      // 3: api.analytics.v1.EventAnalyticsRequest.Event
	(*EventAnalyticsRequest_EventsView)(nil), // 4: api.analytics.v1.EventAnalyticsRequest.EventsView
}
var file_api_analytics_v1_analytics_proto_depIdxs = []int32{
	3, // 0: api.analytics.v1.EventAnalyticsRequest.events:type_name -> api.analytics.v1.EventAnalyticsRequest.Event
	4, // 1: api.analytics.v1.EventAnalyticsRequest.eventsView:type_name -> api.analytics.v1.EventAnalyticsRequest.EventsView
	2, // 2: api.analytics.v1.EventAnalyticsRequest.Event.filters:type_name -> api.analytics.v1.EventAnalyticsRequest.Filter
	2, // 3: api.analytics.v1.EventAnalyticsRequest.EventsView.filters:type_name -> api.analytics.v1.EventAnalyticsRequest.Filter
	0, // 4: api.analytics.v1.Analytics.EventAnalytics:input_type -> api.analytics.v1.EventAnalyticsRequest
	1, // 5: api.analytics.v1.Analytics.EventAnalytics:output_type -> api.analytics.v1.EventAnalyticReply
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_analytics_v1_analytics_proto_init() }
func file_api_analytics_v1_analytics_proto_init() {
	if File_api_analytics_v1_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_analytics_v1_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventAnalyticsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_analytics_v1_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventAnalyticReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_analytics_v1_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventAnalyticsRequest_Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_analytics_v1_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventAnalyticsRequest_Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_analytics_v1_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventAnalyticsRequest_EventsView); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_analytics_v1_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_analytics_v1_analytics_proto_goTypes,
		DependencyIndexes: file_api_analytics_v1_analytics_proto_depIdxs,
		MessageInfos:      file_api_analytics_v1_analytics_proto_msgTypes,
	}.Build()
	File_api_analytics_v1_analytics_proto = out.File
	file_api_analytics_v1_analytics_proto_rawDesc = nil
	file_api_analytics_v1_analytics_proto_goTypes = nil
	file_api_analytics_v1_analytics_proto_depIdxs = nil
}
