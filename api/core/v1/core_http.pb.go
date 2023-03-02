// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.21.7
// source: api/core/v1/core.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCoreDetail = "/api.core.v1.Core/Detail"
const OperationCoreLogin = "/api.core.v1.Core/Login"
const OperationCoreLogout = "/api.core.v1.Core/Logout"
const OperationCoreRegister = "/api.core.v1.Core/Register"
const OperationCoreUnregister = "/api.core.v1.Core/Unregister"
const OperationCoreUpdate = "/api.core.v1.Core/Update"

type CoreHTTPServer interface {
	Detail(context.Context, *emptypb.Empty) (*UserDetailReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *emptypb.Empty) (*LogoutReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Unregister(context.Context, *UnregisterRequest) (*UnregisterReply, error)
	Update(context.Context, *UserInfoUpdateRequest) (*UserInfoUpdateReply, error)
}

func RegisterCoreHTTPServer(s *http.Server, srv CoreHTTPServer) {
	r := s.Route("/")
	r.POST("v1/api/user/register", _Core_Register0_HTTP_Handler(srv))
	r.POST("v1/api/user/login", _Core_Login0_HTTP_Handler(srv))
	r.DELETE("v1/api/user/unregister", _Core_Unregister0_HTTP_Handler(srv))
	r.POST("v1/api/user/logout", _Core_Logout0_HTTP_Handler(srv))
	r.GET("v1/api/user/detail", _Core_Detail0_HTTP_Handler(srv))
	r.PUT("v1/api/user/update", _Core_Update0_HTTP_Handler(srv))
}

func _Core_Register0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Core_Login0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Core_Unregister0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnregisterRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUnregister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Unregister(ctx, req.(*UnregisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UnregisterReply)
		return ctx.Result(200, reply)
	}
}

func _Core_Logout0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _Core_Detail0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detail(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserDetailReply)
		return ctx.Result(200, reply)
	}
}

func _Core_Update0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*UserInfoUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoUpdateReply)
		return ctx.Result(200, reply)
	}
}

type CoreHTTPClient interface {
	Detail(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserDetailReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	Unregister(ctx context.Context, req *UnregisterRequest, opts ...http.CallOption) (rsp *UnregisterReply, err error)
	Update(ctx context.Context, req *UserInfoUpdateRequest, opts ...http.CallOption) (rsp *UserInfoUpdateReply, err error)
}

type CoreHTTPClientImpl struct {
	cc *http.Client
}

func NewCoreHTTPClient(client *http.Client) CoreHTTPClient {
	return &CoreHTTPClientImpl{client}
}

func (c *CoreHTTPClientImpl) Detail(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserDetailReply, error) {
	var out UserDetailReply
	pattern := "v1/api/user/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "v1/api/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) Logout(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "v1/api/user/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "v1/api/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) Unregister(ctx context.Context, in *UnregisterRequest, opts ...http.CallOption) (*UnregisterReply, error) {
	var out UnregisterReply
	pattern := "v1/api/user/unregister"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreUnregister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) Update(ctx context.Context, in *UserInfoUpdateRequest, opts ...http.CallOption) (*UserInfoUpdateReply, error) {
	var out UserInfoUpdateReply
	pattern := "v1/api/user/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
