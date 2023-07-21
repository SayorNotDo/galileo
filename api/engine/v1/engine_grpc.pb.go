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
	Engine_TestEngine_FullMethodName     = "/api.engine.v1.Engine/TestEngine"
	Engine_RunJob_FullMethodName         = "/api.engine.v1.Engine/RunJob"
	Engine_CancelJob_FullMethodName      = "/api.engine.v1.Engine/CancelJob"
	Engine_PauseJob_FullMethodName       = "/api.engine.v1.Engine/PauseJob"
	Engine_ResumeJob_FullMethodName      = "/api.engine.v1.Engine/ResumeJob"
	Engine_DeleteJob_FullMethodName      = "/api.engine.v1.Engine/DeleteJob"
	Engine_AddCronJob_FullMethodName     = "/api.engine.v1.Engine/AddCronJob"
	Engine_UpdateCronJob_FullMethodName  = "/api.engine.v1.Engine/UpdateCronJob"
	Engine_BuildContainer_FullMethodName = "/api.engine.v1.Engine/BuildContainer"
)

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngineClient interface {
	TestEngine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TestEngineReply, error)
	RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (*RunJobReply, error)
	CancelJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PauseJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResumeJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddCronJob(ctx context.Context, in *AddCronJobRequest, opts ...grpc.CallOption) (*AddCronJobReply, error)
	UpdateCronJob(ctx context.Context, in *UpdateCronJobRequest, opts ...grpc.CallOption) (*UpdateCronJobReply, error)
	BuildContainer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BuildContainerReply, error)
}

type engineClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineClient(cc grpc.ClientConnInterface) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) TestEngine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TestEngineReply, error) {
	out := new(TestEngineReply)
	err := c.cc.Invoke(ctx, Engine_TestEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (*RunJobReply, error) {
	out := new(RunJobReply)
	err := c.cc.Invoke(ctx, Engine_RunJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) CancelJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_CancelJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) PauseJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_PauseJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) ResumeJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_ResumeJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) DeleteJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_DeleteJob_FullMethodName, in, out, opts...)
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

func (c *engineClient) UpdateCronJob(ctx context.Context, in *UpdateCronJobRequest, opts ...grpc.CallOption) (*UpdateCronJobReply, error) {
	out := new(UpdateCronJobReply)
	err := c.cc.Invoke(ctx, Engine_UpdateCronJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) BuildContainer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BuildContainerReply, error) {
	out := new(BuildContainerReply)
	err := c.cc.Invoke(ctx, Engine_BuildContainer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
// All implementations must embed UnimplementedEngineServer
// for forward compatibility
type EngineServer interface {
	TestEngine(context.Context, *emptypb.Empty) (*TestEngineReply, error)
	RunJob(context.Context, *RunJobRequest) (*RunJobReply, error)
	CancelJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	PauseJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ResumeJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	DeleteJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	AddCronJob(context.Context, *AddCronJobRequest) (*AddCronJobReply, error)
	UpdateCronJob(context.Context, *UpdateCronJobRequest) (*UpdateCronJobReply, error)
	BuildContainer(context.Context, *emptypb.Empty) (*BuildContainerReply, error)
	mustEmbedUnimplementedEngineServer()
}

// UnimplementedEngineServer must be embedded to have forward compatible implementations.
type UnimplementedEngineServer struct {
}

func (UnimplementedEngineServer) TestEngine(context.Context, *emptypb.Empty) (*TestEngineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestEngine not implemented")
}
func (UnimplementedEngineServer) RunJob(context.Context, *RunJobRequest) (*RunJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunJob not implemented")
}
func (UnimplementedEngineServer) CancelJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJob not implemented")
}
func (UnimplementedEngineServer) PauseJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseJob not implemented")
}
func (UnimplementedEngineServer) ResumeJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeJob not implemented")
}
func (UnimplementedEngineServer) DeleteJob(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (UnimplementedEngineServer) AddCronJob(context.Context, *AddCronJobRequest) (*AddCronJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCronJob not implemented")
}
func (UnimplementedEngineServer) UpdateCronJob(context.Context, *UpdateCronJobRequest) (*UpdateCronJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCronJob not implemented")
}
func (UnimplementedEngineServer) BuildContainer(context.Context, *emptypb.Empty) (*BuildContainerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildContainer not implemented")
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

func _Engine_TestEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).TestEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_TestEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).TestEngine(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
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

func _Engine_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_CancelJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).CancelJob(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_PauseJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).PauseJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_PauseJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).PauseJob(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_ResumeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).ResumeJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_ResumeJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).ResumeJob(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_DeleteJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).DeleteJob(ctx, req.(*emptypb.Empty))
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

func _Engine_BuildContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).BuildContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_BuildContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).BuildContainer(ctx, req.(*emptypb.Empty))
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
			MethodName: "TestEngine",
			Handler:    _Engine_TestEngine_Handler,
		},
		{
			MethodName: "RunJob",
			Handler:    _Engine_RunJob_Handler,
		},
		{
			MethodName: "CancelJob",
			Handler:    _Engine_CancelJob_Handler,
		},
		{
			MethodName: "PauseJob",
			Handler:    _Engine_PauseJob_Handler,
		},
		{
			MethodName: "ResumeJob",
			Handler:    _Engine_ResumeJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _Engine_DeleteJob_Handler,
		},
		{
			MethodName: "AddCronJob",
			Handler:    _Engine_AddCronJob_Handler,
		},
		{
			MethodName: "UpdateCronJob",
			Handler:    _Engine_UpdateCronJob_Handler,
		},
		{
			MethodName: "BuildContainer",
			Handler:    _Engine_BuildContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/engine/v1/engine.proto",
}
