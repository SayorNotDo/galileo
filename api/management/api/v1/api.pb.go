// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.7
// source: api/management/api/v1/api.proto

package api

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

type MethodType int32

const (
	MethodType_GET     MethodType = 0
	MethodType_POST    MethodType = 1
	MethodType_PUT     MethodType = 2
	MethodType_DELETE  MethodType = 3
	MethodType_PATCH   MethodType = 4
	MethodType_OPTIONS MethodType = 5
	MethodType_HEAD    MethodType = 6
)

// Enum value maps for MethodType.
var (
	MethodType_name = map[int32]string{
		0: "GET",
		1: "POST",
		2: "PUT",
		3: "DELETE",
		4: "PATCH",
		5: "OPTIONS",
		6: "HEAD",
	}
	MethodType_value = map[string]int32{
		"GET":     0,
		"POST":    1,
		"PUT":     2,
		"DELETE":  3,
		"PATCH":   4,
		"OPTIONS": 5,
		"HEAD":    6,
	}
)

func (x MethodType) Enum() *MethodType {
	p := new(MethodType)
	*p = x
	return p
}

func (x MethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_management_api_v1_api_proto_enumTypes[0].Descriptor()
}

func (MethodType) Type() protoreflect.EnumType {
	return &file_api_management_api_v1_api_proto_enumTypes[0]
}

func (x MethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodType.Descriptor instead.
func (MethodType) EnumDescriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{0}
}

type DebugApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DebugApiRequest) Reset() {
	*x = DebugApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugApiRequest) ProtoMessage() {}

func (x *DebugApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugApiRequest.ProtoReflect.Descriptor instead.
func (*DebugApiRequest) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{0}
}

type DebugApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DebugApiReply) Reset() {
	*x = DebugApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugApiReply) ProtoMessage() {}

func (x *DebugApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugApiReply.ProtoReflect.Descriptor instead.
func (*DebugApiReply) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{1}
}

type CreateApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url         string     `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Type        MethodType `protobuf:"varint,3,opt,name=type,proto3,enum=api.management.api.MethodType" json:"type,omitempty"`
	Status      int32      `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Body        string     `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	QueryParams string     `protobuf:"bytes,6,opt,name=queryParams,proto3" json:"queryParams,omitempty"`
	Response    string     `protobuf:"bytes,7,opt,name=response,proto3" json:"response,omitempty"`
	Module      string     `protobuf:"bytes,8,opt,name=module,proto3" json:"module,omitempty"`
	Description string     `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateApiRequest) Reset() {
	*x = CreateApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiRequest) ProtoMessage() {}

func (x *CreateApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiRequest.ProtoReflect.Descriptor instead.
func (*CreateApiRequest) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateApiRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateApiRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateApiRequest) GetType() MethodType {
	if x != nil {
		return x.Type
	}
	return MethodType_GET
}

func (x *CreateApiRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateApiRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CreateApiRequest) GetQueryParams() string {
	if x != nil {
		return x.QueryParams
	}
	return ""
}

func (x *CreateApiRequest) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *CreateApiRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *CreateApiRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64 `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *CreateApiReply) Reset() {
	*x = CreateApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiReply) ProtoMessage() {}

func (x *CreateApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiReply.ProtoReflect.Descriptor instead.
func (*CreateApiReply) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateApiReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateApiReply) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type UpdateApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateApiRequest) Reset() {
	*x = UpdateApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiRequest) ProtoMessage() {}

func (x *UpdateApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiRequest.ProtoReflect.Descriptor instead.
func (*UpdateApiRequest) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{4}
}

type UpdateApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateApiReply) Reset() {
	*x = UpdateApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiReply) ProtoMessage() {}

func (x *UpdateApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiReply.ProtoReflect.Descriptor instead.
func (*UpdateApiReply) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{5}
}

type DeleteApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApiRequest) Reset() {
	*x = DeleteApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiRequest) ProtoMessage() {}

func (x *DeleteApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiRequest.ProtoReflect.Descriptor instead.
func (*DeleteApiRequest) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{6}
}

type DeleteApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApiReply) Reset() {
	*x = DeleteApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiReply) ProtoMessage() {}

func (x *DeleteApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiReply.ProtoReflect.Descriptor instead.
func (*DeleteApiReply) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{7}
}

type GetApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetApiRequest) Reset() {
	*x = GetApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiRequest) ProtoMessage() {}

func (x *GetApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiRequest.ProtoReflect.Descriptor instead.
func (*GetApiRequest) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{8}
}

type GetApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetApiReply) Reset() {
	*x = GetApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiReply) ProtoMessage() {}

func (x *GetApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiReply.ProtoReflect.Descriptor instead.
func (*GetApiReply) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{9}
}

type ListApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApiRequest) Reset() {
	*x = ListApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiRequest) ProtoMessage() {}

func (x *ListApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiRequest.ProtoReflect.Descriptor instead.
func (*ListApiRequest) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{10}
}

type ListApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApiReply) Reset() {
	*x = ListApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_management_api_v1_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiReply) ProtoMessage() {}

func (x *ListApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_management_api_v1_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiReply.ProtoReflect.Descriptor instead.
func (*ListApiReply) Descriptor() ([]byte, []int) {
	return file_api_management_api_v1_api_proto_rawDescGZIP(), []int{11}
}

var File_api_management_api_v1_api_proto protoreflect.FileDescriptor

var file_api_management_api_v1_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x62, 0x75, 0x67, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x90, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x56, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x53, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x06,
	0x32, 0xa5, 0x04, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x7d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x69, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x69, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x52, 0x0a, 0x08, 0x44, 0x65, 0x62, 0x75, 0x67, 0x41, 0x70, 0x69, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x70, 0x69, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x36, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x50, 0x01,
	0x5a, 0x1e, 0x67, 0x61, 0x6c, 0x69, 0x6c, 0x65, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_management_api_v1_api_proto_rawDescOnce sync.Once
	file_api_management_api_v1_api_proto_rawDescData = file_api_management_api_v1_api_proto_rawDesc
)

func file_api_management_api_v1_api_proto_rawDescGZIP() []byte {
	file_api_management_api_v1_api_proto_rawDescOnce.Do(func() {
		file_api_management_api_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_management_api_v1_api_proto_rawDescData)
	})
	return file_api_management_api_v1_api_proto_rawDescData
}

var file_api_management_api_v1_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_management_api_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_management_api_v1_api_proto_goTypes = []interface{}{
	(MethodType)(0),          // 0: api.management.api.MethodType
	(*DebugApiRequest)(nil),  // 1: api.management.api.DebugApiRequest
	(*DebugApiReply)(nil),    // 2: api.management.api.DebugApiReply
	(*CreateApiRequest)(nil), // 3: api.management.api.CreateApiRequest
	(*CreateApiReply)(nil),   // 4: api.management.api.CreateApiReply
	(*UpdateApiRequest)(nil), // 5: api.management.api.UpdateApiRequest
	(*UpdateApiReply)(nil),   // 6: api.management.api.UpdateApiReply
	(*DeleteApiRequest)(nil), // 7: api.management.api.DeleteApiRequest
	(*DeleteApiReply)(nil),   // 8: api.management.api.DeleteApiReply
	(*GetApiRequest)(nil),    // 9: api.management.api.GetApiRequest
	(*GetApiReply)(nil),      // 10: api.management.api.GetApiReply
	(*ListApiRequest)(nil),   // 11: api.management.api.ListApiRequest
	(*ListApiReply)(nil),     // 12: api.management.api.ListApiReply
}
var file_api_management_api_v1_api_proto_depIdxs = []int32{
	0,  // 0: api.management.api.CreateApiRequest.type:type_name -> api.management.api.MethodType
	3,  // 1: api.management.api.Api.CreateApi:input_type -> api.management.api.CreateApiRequest
	5,  // 2: api.management.api.Api.UpdateApi:input_type -> api.management.api.UpdateApiRequest
	7,  // 3: api.management.api.Api.DeleteApi:input_type -> api.management.api.DeleteApiRequest
	9,  // 4: api.management.api.Api.GetApi:input_type -> api.management.api.GetApiRequest
	1,  // 5: api.management.api.Api.DebugApi:input_type -> api.management.api.DebugApiRequest
	11, // 6: api.management.api.Api.ListApi:input_type -> api.management.api.ListApiRequest
	4,  // 7: api.management.api.Api.CreateApi:output_type -> api.management.api.CreateApiReply
	6,  // 8: api.management.api.Api.UpdateApi:output_type -> api.management.api.UpdateApiReply
	8,  // 9: api.management.api.Api.DeleteApi:output_type -> api.management.api.DeleteApiReply
	10, // 10: api.management.api.Api.GetApi:output_type -> api.management.api.GetApiReply
	2,  // 11: api.management.api.Api.DebugApi:output_type -> api.management.api.DebugApiReply
	12, // 12: api.management.api.Api.ListApi:output_type -> api.management.api.ListApiReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_management_api_v1_api_proto_init() }
func file_api_management_api_v1_api_proto_init() {
	if File_api_management_api_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_management_api_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugApiRequest); i {
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
		file_api_management_api_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugApiReply); i {
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
		file_api_management_api_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiRequest); i {
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
		file_api_management_api_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiReply); i {
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
		file_api_management_api_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApiRequest); i {
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
		file_api_management_api_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApiReply); i {
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
		file_api_management_api_v1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiRequest); i {
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
		file_api_management_api_v1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiReply); i {
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
		file_api_management_api_v1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiRequest); i {
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
		file_api_management_api_v1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiReply); i {
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
		file_api_management_api_v1_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiRequest); i {
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
		file_api_management_api_v1_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiReply); i {
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
			RawDescriptor: file_api_management_api_v1_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_management_api_v1_api_proto_goTypes,
		DependencyIndexes: file_api_management_api_v1_api_proto_depIdxs,
		EnumInfos:         file_api_management_api_v1_api_proto_enumTypes,
		MessageInfos:      file_api_management_api_v1_api_proto_msgTypes,
	}.Build()
	File_api_management_api_v1_api_proto = out.File
	file_api_management_api_v1_api_proto_rawDesc = nil
	file_api_management_api_v1_api_proto_goTypes = nil
	file_api_management_api_v1_api_proto_depIdxs = nil
}