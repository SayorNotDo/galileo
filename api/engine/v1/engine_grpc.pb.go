// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/engine/v1/engine.proto

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
	Engine_RunJob_FullMethodName           = "/api.engine.v1.Engine/RunJob"
	Engine_AddCronJob_FullMethodName       = "/api.engine.v1.Engine/AddCronJob"
	Engine_AddDelayedJob_FullMethodName    = "/api.engine.v1.Engine/AddDelayedJob"
	Engine_UpdateCronJob_FullMethodName    = "/api.engine.v1.Engine/UpdateCronJob"
	Engine_CreateContainer_FullMethodName  = "/api.engine.v1.Engine/CreateContainer"
	Engine_ListContainers_FullMethodName   = "/api.engine.v1.Engine/ListContainers"
	Engine_InspectContainer_FullMethodName = "/api.engine.v1.Engine/InspectContainer"
)

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngineClient interface {
	RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (*RunJobReply, error)
	AddCronJob(ctx context.Context, in *AddCronJobRequest, opts ...grpc.CallOption) (*AddCronJobReply, error)
	AddDelayedJob(ctx context.Context, in *AddDelayedJobRequest, opts ...grpc.CallOption) (*AddDelayedJobReply, error)
	UpdateCronJob(ctx context.Context, in *UpdateCronJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerReply, error)
	ListContainers(ctx context.Context, in *ListContainerRequest, opts ...grpc.CallOption) (*ListContainersReply, error)
	InspectContainer(ctx context.Context, in *InspectContainerRequest, opts ...grpc.CallOption) (*InspectContainerReply, error)
}

type engineClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineClient(cc grpc.ClientConnInterface) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (*RunJobReply, error) {
	out := new(RunJobReply)
	err := c.cc.Invoke(ctx, Engine_RunJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) AddCronJob(ctx context.Context, in *AddCronJobRequest, opts ...grpc.CallOption) (*AddCronJobReply, error) {
	out := new(AddCronJobReply)
	err := c.cc.Invoke(ctx, Engine_AddCronJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) AddDelayedJob(ctx context.Context, in *AddDelayedJobRequest, opts ...grpc.CallOption) (*AddDelayedJobReply, error) {
	out := new(AddDelayedJobReply)
	err := c.cc.Invoke(ctx, Engine_AddDelayedJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) UpdateCronJob(ctx context.Context, in *UpdateCronJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_UpdateCronJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerReply, error) {
	out := new(CreateContainerReply)
	err := c.cc.Invoke(ctx, Engine_CreateContainer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) ListContainers(ctx context.Context, in *ListContainerRequest, opts ...grpc.CallOption) (*ListContainersReply, error) {
	out := new(ListContainersReply)
	err := c.cc.Invoke(ctx, Engine_ListContainers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) InspectContainer(ctx context.Context, in *InspectContainerRequest, opts ...grpc.CallOption) (*InspectContainerReply, error) {
	out := new(InspectContainerReply)
	err := c.cc.Invoke(ctx, Engine_InspectContainer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
// All implementations must embed UnimplementedEngineServer
// for forward compatibility
type EngineServer interface {
	RunJob(context.Context, *RunJobRequest) (*RunJobReply, error)
	AddCronJob(context.Context, *AddCronJobRequest) (*AddCronJobReply, error)
	AddDelayedJob(context.Context, *AddDelayedJobRequest) (*AddDelayedJobReply, error)
	UpdateCronJob(context.Context, *UpdateCronJobRequest) (*emptypb.Empty, error)
	CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerReply, error)
	ListContainers(context.Context, *ListContainerRequest) (*ListContainersReply, error)
	InspectContainer(context.Context, *InspectContainerRequest) (*InspectContainerReply, error)
	mustEmbedUnimplementedEngineServer()
}

// UnimplementedEngineServer must be embedded to have forward compatible implementations.
type UnimplementedEngineServer struct {
}

func (UnimplementedEngineServer) RunJob(context.Context, *RunJobRequest) (*RunJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunJob not implemented")
}
func (UnimplementedEngineServer) AddCronJob(context.Context, *AddCronJobRequest) (*AddCronJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCronJob not implemented")
}
func (UnimplementedEngineServer) AddDelayedJob(context.Context, *AddDelayedJobRequest) (*AddDelayedJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDelayedJob not implemented")
}
func (UnimplementedEngineServer) UpdateCronJob(context.Context, *UpdateCronJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCronJob not implemented")
}
func (UnimplementedEngineServer) CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}
func (UnimplementedEngineServer) ListContainers(context.Context, *ListContainerRequest) (*ListContainersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContainers not implemented")
}
func (UnimplementedEngineServer) InspectContainer(context.Context, *InspectContainerRequest) (*InspectContainerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InspectContainer not implemented")
}
func (UnimplementedEngineServer) mustEmbedUnimplementedEngineServer() {}

// UnsafeEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EngineServer will
// result in compilation errors.
type UnsafeEngineServer interface {
	mustEmbedUnimplementedEngineServer()
}

func RegisterEngineServer(s grpc.ServiceRegistrar, srv EngineServer) {
	s.RegisterService(&Engine_ServiceDesc, srv)
}

func _Engine_RunJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).RunJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_RunJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).RunJob(ctx, req.(*RunJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_AddCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).AddCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_AddCronJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).AddCronJob(ctx, req.(*AddCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_AddDelayedJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDelayedJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).AddDelayedJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_AddDelayedJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).AddDelayedJob(ctx, req.(*AddDelayedJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_UpdateCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).UpdateCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_UpdateCronJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).UpdateCronJob(ctx, req.(*UpdateCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).CreateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_CreateContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).CreateContainer(ctx, req.(*CreateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_ListContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).ListContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_ListContainers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).ListContainers(ctx, req.(*ListContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_InspectContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).InspectContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_InspectContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).InspectContainer(ctx, req.(*InspectContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Engine_ServiceDesc is the grpc.ServiceDesc for Engine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Engine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.engine.v1.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunJob",
			Handler:    _Engine_RunJob_Handler,
		},
		{
			MethodName: "AddCronJob",
			Handler:    _Engine_AddCronJob_Handler,
		},
		{
			MethodName: "AddDelayedJob",
			Handler:    _Engine_AddDelayedJob_Handler,
		},
		{
			MethodName: "UpdateCronJob",
			Handler:    _Engine_UpdateCronJob_Handler,
		},
		{
			MethodName: "CreateContainer",
			Handler:    _Engine_CreateContainer_Handler,
		},
		{
			MethodName: "ListContainers",
			Handler:    _Engine_ListContainers_Handler,
		},
		{
			MethodName: "InspectContainer",
			Handler:    _Engine_InspectContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/engine/v1/engine.proto",
}
