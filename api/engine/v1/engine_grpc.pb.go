// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.7
// source: api/engine/v1/engine.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Engine_CreateEngine_FullMethodName = "/api.engine.v1.Engine/CreateEngine"
	Engine_UpdateEngine_FullMethodName = "/api.engine.v1.Engine/UpdateEngine"
	Engine_DeleteEngine_FullMethodName = "/api.engine.v1.Engine/DeleteEngine"
	Engine_GetEngine_FullMethodName    = "/api.engine.v1.Engine/GetEngine"
	Engine_ListEngine_FullMethodName   = "/api.engine.v1.Engine/ListEngine"
)

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngineClient interface {
	CreateEngine(ctx context.Context, in *CreateEngineRequest, opts ...grpc.CallOption) (*CreateEngineReply, error)
	UpdateEngine(ctx context.Context, in *UpdateEngineRequest, opts ...grpc.CallOption) (*UpdateEngineReply, error)
	DeleteEngine(ctx context.Context, in *DeleteEngineRequest, opts ...grpc.CallOption) (*DeleteEngineReply, error)
	GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineReply, error)
	ListEngine(ctx context.Context, in *ListEngineRequest, opts ...grpc.CallOption) (*ListEngineReply, error)
}

type engineClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineClient(cc grpc.ClientConnInterface) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) CreateEngine(ctx context.Context, in *CreateEngineRequest, opts ...grpc.CallOption) (*CreateEngineReply, error) {
	out := new(CreateEngineReply)
	err := c.cc.Invoke(ctx, Engine_CreateEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) UpdateEngine(ctx context.Context, in *UpdateEngineRequest, opts ...grpc.CallOption) (*UpdateEngineReply, error) {
	out := new(UpdateEngineReply)
	err := c.cc.Invoke(ctx, Engine_UpdateEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) DeleteEngine(ctx context.Context, in *DeleteEngineRequest, opts ...grpc.CallOption) (*DeleteEngineReply, error) {
	out := new(DeleteEngineReply)
	err := c.cc.Invoke(ctx, Engine_DeleteEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineReply, error) {
	out := new(GetEngineReply)
	err := c.cc.Invoke(ctx, Engine_GetEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) ListEngine(ctx context.Context, in *ListEngineRequest, opts ...grpc.CallOption) (*ListEngineReply, error) {
	out := new(ListEngineReply)
	err := c.cc.Invoke(ctx, Engine_ListEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
// All implementations must embed UnimplementedEngineServer
// for forward compatibility
type EngineServer interface {
	CreateEngine(context.Context, *CreateEngineRequest) (*CreateEngineReply, error)
	UpdateEngine(context.Context, *UpdateEngineRequest) (*UpdateEngineReply, error)
	DeleteEngine(context.Context, *DeleteEngineRequest) (*DeleteEngineReply, error)
	GetEngine(context.Context, *GetEngineRequest) (*GetEngineReply, error)
	ListEngine(context.Context, *ListEngineRequest) (*ListEngineReply, error)
	mustEmbedUnimplementedEngineServer()
}

// UnimplementedEngineServer must be embedded to have forward compatible implementations.
type UnimplementedEngineServer struct {
}

func (UnimplementedEngineServer) CreateEngine(context.Context, *CreateEngineRequest) (*CreateEngineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEngine not implemented")
}
func (UnimplementedEngineServer) UpdateEngine(context.Context, *UpdateEngineRequest) (*UpdateEngineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEngine not implemented")
}
func (UnimplementedEngineServer) DeleteEngine(context.Context, *DeleteEngineRequest) (*DeleteEngineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEngine not implemented")
}
func (UnimplementedEngineServer) GetEngine(context.Context, *GetEngineRequest) (*GetEngineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEngine not implemented")
}
func (UnimplementedEngineServer) ListEngine(context.Context, *ListEngineRequest) (*ListEngineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEngine not implemented")
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

func _Engine_CreateEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).CreateEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_CreateEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).CreateEngine(ctx, req.(*CreateEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_UpdateEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).UpdateEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_UpdateEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).UpdateEngine(ctx, req.(*UpdateEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_DeleteEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).DeleteEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_DeleteEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).DeleteEngine(ctx, req.(*DeleteEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_GetEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).GetEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_GetEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).GetEngine(ctx, req.(*GetEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_ListEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).ListEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_ListEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).ListEngine(ctx, req.(*ListEngineRequest))
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
			MethodName: "CreateEngine",
			Handler:    _Engine_CreateEngine_Handler,
		},
		{
			MethodName: "UpdateEngine",
			Handler:    _Engine_UpdateEngine_Handler,
		},
		{
			MethodName: "DeleteEngine",
			Handler:    _Engine_DeleteEngine_Handler,
		},
		{
			MethodName: "GetEngine",
			Handler:    _Engine_GetEngine_Handler,
		},
		{
			MethodName: "ListEngine",
			Handler:    _Engine_ListEngine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/engine/v1/engine.proto",
}