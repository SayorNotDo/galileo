// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/core/v1/core.proto

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
	Core_Register_FullMethodName              = "/api.core.v1.Core/Register"
	Core_Login_FullMethodName                 = "/api.core.v1.Core/Login"
	Core_Logout_FullMethodName                = "/api.core.v1.Core/Logout"
	Core_UserInfo_FullMethodName              = "/api.core.v1.Core/UserInfo"
	Core_CurrentUserInfo_FullMethodName       = "/api.core.v1.Core/CurrentUserInfo"
	Core_UpdatePassword_FullMethodName        = "/api.core.v1.Core/UpdatePassword"
	Core_ResetPassword_FullMethodName         = "/api.core.v1.Core/ResetPassword"
	Core_DeleteUser_FullMethodName            = "/api.core.v1.Core/DeleteUser"
	Core_ListUsers_FullMethodName             = "/api.core.v1.Core/ListUsers"
	Core_UserGroups_FullMethodName            = "/api.core.v1.Core/UserGroups"
	Core_ListUserGroups_FullMethodName        = "/api.core.v1.Core/ListUserGroups"
	Core_GetUserProjectList_FullMethodName    = "/api.core.v1.Core/GetUserProjectList"
	Core_GetUserLatestActivity_FullMethodName = "/api.core.v1.Core/GetUserLatestActivity"
	Core_TrackReportData_FullMethodName       = "/api.core.v1.Core/TrackReportData"
	Core_ExecuteToken_FullMethodName          = "/api.core.v1.Core/ExecuteToken"
	Core_InspectContainer_FullMethodName      = "/api.core.v1.Core/InspectContainer"
)

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	CurrentUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserInfoReply, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResetPassword(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResetPasswordReply, error)
	DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListUsers(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
	UserGroups(ctx context.Context, in *GroupInfoRequest, opts ...grpc.CallOption) (*GroupInfo, error)
	ListUserGroups(ctx context.Context, in *ListUserGroupsRequest, opts ...grpc.CallOption) (*UserGroupListReply, error)
	GetUserProjectList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserProjectListReply, error)
	GetUserLatestActivity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserLatestActivityReply, error)
	TrackReportData(ctx context.Context, in *TrackReportDataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ExecuteToken(ctx context.Context, in *ExecuteTokenRequest, opts ...grpc.CallOption) (*ExecuteTokenReply, error)
	InspectContainer(ctx context.Context, in *InspectContainerRequest, opts ...grpc.CallOption) (*ContainerInfo, error)
}

type coreClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreClient(cc grpc.ClientConnInterface) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, Core_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Core_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Core_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, Core_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CurrentUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, Core_CurrentUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Core_UpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ResetPassword(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResetPasswordReply, error) {
	out := new(ResetPasswordReply)
	err := c.cc.Invoke(ctx, Core_ResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Core_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListUsers(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, Core_ListUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) UserGroups(ctx context.Context, in *GroupInfoRequest, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, Core_UserGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListUserGroups(ctx context.Context, in *ListUserGroupsRequest, opts ...grpc.CallOption) (*UserGroupListReply, error) {
	out := new(UserGroupListReply)
	err := c.cc.Invoke(ctx, Core_ListUserGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetUserProjectList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserProjectListReply, error) {
	out := new(UserProjectListReply)
	err := c.cc.Invoke(ctx, Core_GetUserProjectList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetUserLatestActivity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserLatestActivityReply, error) {
	out := new(UserLatestActivityReply)
	err := c.cc.Invoke(ctx, Core_GetUserLatestActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) TrackReportData(ctx context.Context, in *TrackReportDataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Core_TrackReportData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ExecuteToken(ctx context.Context, in *ExecuteTokenRequest, opts ...grpc.CallOption) (*ExecuteTokenReply, error) {
	out := new(ExecuteTokenReply)
	err := c.cc.Invoke(ctx, Core_ExecuteToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) InspectContainer(ctx context.Context, in *InspectContainerRequest, opts ...grpc.CallOption) (*ContainerInfo, error) {
	out := new(ContainerInfo)
	err := c.cc.Invoke(ctx, Core_InspectContainer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
// All implementations must embed UnimplementedCoreServer
// for forward compatibility
type CoreServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	CurrentUserInfo(context.Context, *emptypb.Empty) (*UserInfoReply, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error)
	ResetPassword(context.Context, *emptypb.Empty) (*ResetPasswordReply, error)
	DeleteUser(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	ListUsers(context.Context, *ListUserRequest) (*ListUserReply, error)
	UserGroups(context.Context, *GroupInfoRequest) (*GroupInfo, error)
	ListUserGroups(context.Context, *ListUserGroupsRequest) (*UserGroupListReply, error)
	GetUserProjectList(context.Context, *emptypb.Empty) (*UserProjectListReply, error)
	GetUserLatestActivity(context.Context, *emptypb.Empty) (*UserLatestActivityReply, error)
	TrackReportData(context.Context, *TrackReportDataRequest) (*emptypb.Empty, error)
	ExecuteToken(context.Context, *ExecuteTokenRequest) (*ExecuteTokenReply, error)
	InspectContainer(context.Context, *InspectContainerRequest) (*ContainerInfo, error)
	mustEmbedUnimplementedCoreServer()
}

// UnimplementedCoreServer must be embedded to have forward compatible implementations.
type UnimplementedCoreServer struct {
}

func (UnimplementedCoreServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedCoreServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCoreServer) Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedCoreServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedCoreServer) CurrentUserInfo(context.Context, *emptypb.Empty) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentUserInfo not implemented")
}
func (UnimplementedCoreServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedCoreServer) ResetPassword(context.Context, *emptypb.Empty) (*ResetPasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedCoreServer) DeleteUser(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedCoreServer) ListUsers(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedCoreServer) UserGroups(context.Context, *GroupInfoRequest) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGroups not implemented")
}
func (UnimplementedCoreServer) ListUserGroups(context.Context, *ListUserGroupsRequest) (*UserGroupListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserGroups not implemented")
}
func (UnimplementedCoreServer) GetUserProjectList(context.Context, *emptypb.Empty) (*UserProjectListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProjectList not implemented")
}
func (UnimplementedCoreServer) GetUserLatestActivity(context.Context, *emptypb.Empty) (*UserLatestActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLatestActivity not implemented")
}
func (UnimplementedCoreServer) TrackReportData(context.Context, *TrackReportDataRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackReportData not implemented")
}
func (UnimplementedCoreServer) ExecuteToken(context.Context, *ExecuteTokenRequest) (*ExecuteTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteToken not implemented")
}
func (UnimplementedCoreServer) InspectContainer(context.Context, *InspectContainerRequest) (*ContainerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InspectContainer not implemented")
}
func (UnimplementedCoreServer) mustEmbedUnimplementedCoreServer() {}

// UnsafeCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServer will
// result in compilation errors.
type UnsafeCoreServer interface {
	mustEmbedUnimplementedCoreServer()
}

func RegisterCoreServer(s grpc.ServiceRegistrar, srv CoreServer) {
	s.RegisterService(&Core_ServiceDesc, srv)
}

func _Core_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Logout(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CurrentUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CurrentUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_CurrentUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CurrentUserInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ResetPassword(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteUser(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListUsers(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_UserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).UserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_UserGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).UserGroups(ctx, req.(*GroupInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_ListUserGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListUserGroups(ctx, req.(*ListUserGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetUserProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetUserProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_GetUserProjectList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetUserProjectList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetUserLatestActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetUserLatestActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_GetUserLatestActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetUserLatestActivity(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_TrackReportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackReportDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).TrackReportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_TrackReportData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).TrackReportData(ctx, req.(*TrackReportDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ExecuteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ExecuteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_ExecuteToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ExecuteToken(ctx, req.(*ExecuteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_InspectContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).InspectContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_InspectContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).InspectContainer(ctx, req.(*InspectContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Core_ServiceDesc is the grpc.ServiceDesc for Core service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Core_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.core.v1.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Core_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Core_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Core_Logout_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Core_UserInfo_Handler,
		},
		{
			MethodName: "CurrentUserInfo",
			Handler:    _Core_CurrentUserInfo_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Core_UpdatePassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Core_ResetPassword_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Core_DeleteUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Core_ListUsers_Handler,
		},
		{
			MethodName: "UserGroups",
			Handler:    _Core_UserGroups_Handler,
		},
		{
			MethodName: "ListUserGroups",
			Handler:    _Core_ListUserGroups_Handler,
		},
		{
			MethodName: "GetUserProjectList",
			Handler:    _Core_GetUserProjectList_Handler,
		},
		{
			MethodName: "GetUserLatestActivity",
			Handler:    _Core_GetUserLatestActivity_Handler,
		},
		{
			MethodName: "TrackReportData",
			Handler:    _Core_TrackReportData_Handler,
		},
		{
			MethodName: "ExecuteToken",
			Handler:    _Core_ExecuteToken_Handler,
		},
		{
			MethodName: "InspectContainer",
			Handler:    _Core_InspectContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/core/v1/core.proto",
}
