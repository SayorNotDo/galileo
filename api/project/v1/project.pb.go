// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: api/project/v1/project.proto

package v1

import (
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

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Remark     string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *CreateProjectRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type CreateProjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ProjectInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateProjectReply) Reset() {
	*x = CreateProjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectReply) ProtoMessage() {}

func (x *CreateProjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectReply.ProtoReflect.Descriptor instead.
func (*CreateProjectReply) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectReply) GetData() *ProjectInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Remark     string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Id         uint32 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Status     uint32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProjectInfo) Reset() {
	*x = ProjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInfo) ProtoMessage() {}

func (x *ProjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInfo.ProtoReflect.Descriptor instead.
func (*ProjectInfo) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ProjectInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProjectInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectInfo) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProjectRequest) Reset() {
	*x = UpdateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectRequest) ProtoMessage() {}

func (x *UpdateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{3}
}

type UpdateProjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProjectReply) Reset() {
	*x = UpdateProjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectReply) ProtoMessage() {}

func (x *UpdateProjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectReply.ProtoReflect.Descriptor instead.
func (*UpdateProjectReply) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{4}
}

type DeleteProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProjectRequest) Reset() {
	*x = DeleteProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectRequest) ProtoMessage() {}

func (x *DeleteProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{5}
}

type DeleteProjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProjectReply) Reset() {
	*x = DeleteProjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectReply) ProtoMessage() {}

func (x *DeleteProjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectReply.ProtoReflect.Descriptor instead.
func (*DeleteProjectReply) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{6}
}

type GetProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProjectRequest) Reset() {
	*x = GetProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectRequest) ProtoMessage() {}

func (x *GetProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectRequest.ProtoReflect.Descriptor instead.
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{7}
}

type GetProjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProjectReply) Reset() {
	*x = GetProjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectReply) ProtoMessage() {}

func (x *GetProjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectReply.ProtoReflect.Descriptor instead.
func (*GetProjectReply) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{8}
}

type ListProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectRequest) Reset() {
	*x = ListProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectRequest) ProtoMessage() {}

func (x *ListProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectRequest.ProtoReflect.Descriptor instead.
func (*ListProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{9}
}

type ListProjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectReply) Reset() {
	*x = ListProjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v1_project_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectReply) ProtoMessage() {}

func (x *ListProjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v1_project_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectReply.ProtoReflect.Descriptor instead.
func (*ListProjectReply) Descriptor() ([]byte, []int) {
	return file_api_project_v1_project_proto_rawDescGZIP(), []int{10}
}

var File_api_project_v1_project_proto protoreflect.FileDescriptor

var file_api_project_v1_project_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x62,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x16, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc1, 0x03, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x59, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x2d, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x19, 0x67, 0x61, 0x6c, 0x69, 0x6c, 0x65, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_project_v1_project_proto_rawDescOnce sync.Once
	file_api_project_v1_project_proto_rawDescData = file_api_project_v1_project_proto_rawDesc
)

func file_api_project_v1_project_proto_rawDescGZIP() []byte {
	file_api_project_v1_project_proto_rawDescOnce.Do(func() {
		file_api_project_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_project_v1_project_proto_rawDescData)
	})
	return file_api_project_v1_project_proto_rawDescData
}

var file_api_project_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_project_v1_project_proto_goTypes = []interface{}{
	(*CreateProjectRequest)(nil), // 0: api.project.v1.CreateProjectRequest
	(*CreateProjectReply)(nil),   // 1: api.project.v1.CreateProjectReply
	(*ProjectInfo)(nil),          // 2: api.project.v1.ProjectInfo
	(*UpdateProjectRequest)(nil), // 3: api.project.v1.UpdateProjectRequest
	(*UpdateProjectReply)(nil),   // 4: api.project.v1.UpdateProjectReply
	(*DeleteProjectRequest)(nil), // 5: api.project.v1.DeleteProjectRequest
	(*DeleteProjectReply)(nil),   // 6: api.project.v1.DeleteProjectReply
	(*GetProjectRequest)(nil),    // 7: api.project.v1.GetProjectRequest
	(*GetProjectReply)(nil),      // 8: api.project.v1.GetProjectReply
	(*ListProjectRequest)(nil),   // 9: api.project.v1.ListProjectRequest
	(*ListProjectReply)(nil),     // 10: api.project.v1.ListProjectReply
}
var file_api_project_v1_project_proto_depIdxs = []int32{
	2,  // 0: api.project.v1.CreateProjectReply.data:type_name -> api.project.v1.ProjectInfo
	0,  // 1: api.project.v1.Project.CreateProject:input_type -> api.project.v1.CreateProjectRequest
	3,  // 2: api.project.v1.Project.UpdateProject:input_type -> api.project.v1.UpdateProjectRequest
	5,  // 3: api.project.v1.Project.DeleteProject:input_type -> api.project.v1.DeleteProjectRequest
	7,  // 4: api.project.v1.Project.GetProject:input_type -> api.project.v1.GetProjectRequest
	9,  // 5: api.project.v1.Project.ListProject:input_type -> api.project.v1.ListProjectRequest
	1,  // 6: api.project.v1.Project.CreateProject:output_type -> api.project.v1.CreateProjectReply
	4,  // 7: api.project.v1.Project.UpdateProject:output_type -> api.project.v1.UpdateProjectReply
	6,  // 8: api.project.v1.Project.DeleteProject:output_type -> api.project.v1.DeleteProjectReply
	8,  // 9: api.project.v1.Project.GetProject:output_type -> api.project.v1.GetProjectReply
	10, // 10: api.project.v1.Project.ListProject:output_type -> api.project.v1.ListProjectReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_project_v1_project_proto_init() }
func file_api_project_v1_project_proto_init() {
	if File_api_project_v1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_project_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
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
		file_api_project_v1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectReply); i {
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
		file_api_project_v1_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInfo); i {
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
		file_api_project_v1_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProjectRequest); i {
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
		file_api_project_v1_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProjectReply); i {
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
		file_api_project_v1_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectRequest); i {
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
		file_api_project_v1_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectReply); i {
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
		file_api_project_v1_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectRequest); i {
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
		file_api_project_v1_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectReply); i {
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
		file_api_project_v1_project_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectRequest); i {
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
		file_api_project_v1_project_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectReply); i {
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
			RawDescriptor: file_api_project_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_project_v1_project_proto_goTypes,
		DependencyIndexes: file_api_project_v1_project_proto_depIdxs,
		MessageInfos:      file_api_project_v1_project_proto_msgTypes,
	}.Build()
	File_api_project_v1_project_proto = out.File
	file_api_project_v1_project_proto_rawDesc = nil
	file_api_project_v1_project_proto_goTypes = nil
	file_api_project_v1_project_proto_depIdxs = nil
}
