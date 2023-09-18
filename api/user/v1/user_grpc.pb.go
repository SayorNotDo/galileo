// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/user/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	User_CreateUser_FullMethodName       = "/api.user.v1.User/CreateUser"
	User_UpdateUserInfo_FullMethodName   = "/api.user.v1.User/UpdateUserInfo"
	User_GetUserInfo_FullMethodName      = "/api.user.v1.User/GetUserInfo"
	User_ListUser_FullMethodName         = "/api.user.v1.User/ListUser"
	User_ValidateUser_FullMethodName     = "/api.user.v1.User/ValidateUser"
	User_ResetPassword_FullMethodName    = "/api.user.v1.User/ResetPassword"
	User_UpdatePassword_FullMethodName   = "/api.user.v1.User/UpdatePassword"
	User_SoftDeleteUser_FullMethodName   = "/api.user.v1.User/SoftDeleteUser"
	User_SetToken_FullMethodName         = "/api.user.v1.User/SetToken"
	User_EmptyToken_FullMethodName       = "/api.user.v1.User/EmptyToken"
	User_GetUserGroupList_FullMethodName = "/api.user.v1.User/GetUserGroupList"
	User_GetUserGroup_FullMethodName     = "/api.user.v1.User/GetUserGroup"
	User_UpdateUserGroup_FullMethodName  = "/api.user.v1.User/UpdateUserGroup"
	User_CreateUserGroup_FullMethodName  = "/api.user.v1.User/CreateUserGroup"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error)
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
	ValidateUser(ctx context.Context, in *ValidateUserRequest, opts ...grpc.CallOption) (*ValidateUserReply, error)
	ResetPassword(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SoftDeleteUser(ctx context.Context, in *SoftDeleteRequest, opts ...grpc.CallOption) (*SoftDeleteReply, error)
	SetToken(ctx context.Context, in *SetTokenRequest, opts ...grpc.CallOption) (*SetTokenReply, error)
	EmptyToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EmptyTokenReply, error)
	GetUserGroupList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserGroupListReply, error)
	GetUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*GroupInfo, error)
	UpdateUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*GroupInfo, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, User_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserInfo(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_UpdateUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, User_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, User_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ValidateUser(ctx context.Context, in *ValidateUserRequest, opts ...grpc.CallOption) (*ValidateUserReply, error) {
	out := new(ValidateUserReply)
	err := c.cc.Invoke(ctx, User_ValidateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ResetPassword(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_ResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_UpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SoftDeleteUser(ctx context.Context, in *SoftDeleteRequest, opts ...grpc.CallOption) (*SoftDeleteReply, error) {
	out := new(SoftDeleteReply)
	err := c.cc.Invoke(ctx, User_SoftDeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetToken(ctx context.Context, in *SetTokenRequest, opts ...grpc.CallOption) (*SetTokenReply, error) {
	out := new(SetTokenReply)
	err := c.cc.Invoke(ctx, User_SetToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) EmptyToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EmptyTokenReply, error) {
	out := new(EmptyTokenReply)
	err := c.cc.Invoke(ctx, User_EmptyToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserGroupList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserGroupListReply, error) {
	out := new(UserGroupListReply)
	err := c.cc.Invoke(ctx, User_GetUserGroupList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, User_GetUserGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_UpdateUserGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, User_CreateUserGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserInfo, error)
	UpdateUserInfo(context.Context, *UpdateUserRequest) (*emptypb.Empty, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*UserInfo, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	ValidateUser(context.Context, *ValidateUserRequest) (*ValidateUserReply, error)
	ResetPassword(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error)
	SoftDeleteUser(context.Context, *SoftDeleteRequest) (*SoftDeleteReply, error)
	SetToken(context.Context, *SetTokenRequest) (*SetTokenReply, error)
	EmptyToken(context.Context, *emptypb.Empty) (*EmptyTokenReply, error)
	GetUserGroupList(context.Context, *emptypb.Empty) (*UserGroupListReply, error)
	GetUserGroup(context.Context, *UserGroupRequest) (*GroupInfo, error)
	UpdateUserGroup(context.Context, *UserGroupRequest) (*emptypb.Empty, error)
	CreateUserGroup(context.Context, *UserGroupRequest) (*GroupInfo, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) UpdateUserInfo(context.Context, *UpdateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServer) ListUser(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServer) ValidateUser(context.Context, *ValidateUserRequest) (*ValidateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUser not implemented")
}
func (UnimplementedUserServer) ResetPassword(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServer) SoftDeleteUser(context.Context, *SoftDeleteRequest) (*SoftDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDeleteUser not implemented")
}
func (UnimplementedUserServer) SetToken(context.Context, *SetTokenRequest) (*SetTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToken not implemented")
}
func (UnimplementedUserServer) EmptyToken(context.Context, *emptypb.Empty) (*EmptyTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyToken not implemented")
}
func (UnimplementedUserServer) GetUserGroupList(context.Context, *emptypb.Empty) (*UserGroupListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroupList not implemented")
}
func (UnimplementedUserServer) GetUserGroup(context.Context, *UserGroupRequest) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroup not implemented")
}
func (UnimplementedUserServer) UpdateUserGroup(context.Context, *UserGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGroup not implemented")
}
func (UnimplementedUserServer) CreateUserGroup(context.Context, *UserGroupRequest) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserGroup not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserInfo(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ValidateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ValidateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ValidateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ValidateUser(ctx, req.(*ValidateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ResetPassword(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SoftDeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SoftDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SoftDeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SoftDeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SoftDeleteUser(ctx, req.(*SoftDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetToken(ctx, req.(*SetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_EmptyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EmptyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_EmptyToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EmptyToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserGroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserGroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserGroupList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserGroupList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserGroup(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserGroup(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUserGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUserGroup(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _User_UpdateUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _User_ListUser_Handler,
		},
		{
			MethodName: "ValidateUser",
			Handler:    _User_ValidateUser_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _User_ResetPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _User_UpdatePassword_Handler,
		},
		{
			MethodName: "SoftDeleteUser",
			Handler:    _User_SoftDeleteUser_Handler,
		},
		{
			MethodName: "SetToken",
			Handler:    _User_SetToken_Handler,
		},
		{
			MethodName: "EmptyToken",
			Handler:    _User_EmptyToken_Handler,
		},
		{
			MethodName: "GetUserGroupList",
			Handler:    _User_GetUserGroupList_Handler,
		},
		{
			MethodName: "GetUserGroup",
			Handler:    _User_GetUserGroup_Handler,
		},
		{
			MethodName: "UpdateUserGroup",
			Handler:    _User_UpdateUserGroup_Handler,
		},
		{
			MethodName: "CreateUserGroup",
			Handler:    _User_CreateUserGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user.proto",
}
