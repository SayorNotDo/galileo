// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.7
// source: api/management/v1/management.proto

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

const OperationManagementBaseInfo = "/api.management.v1.Management/BaseInfo"

type ManagementHTTPServer interface {
	BaseInfo(context.Context, *emptypb.Empty) (*BaseInfoReply, error)
}

func RegisterManagementHTTPServer(s *http.Server, srv ManagementHTTPServer) {
	r := s.Route("/")
	r.GET("v1/api/management/info", _Management_BaseInfo0_HTTP_Handler(srv))
}

func _Management_BaseInfo0_HTTP_Handler(srv ManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationManagementBaseInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BaseInfo(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BaseInfoReply)
		return ctx.Result(200, reply)
	}
}

type ManagementHTTPClient interface {
	BaseInfo(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *BaseInfoReply, err error)
}

type ManagementHTTPClientImpl struct {
	cc *http.Client
}

func NewManagementHTTPClient(client *http.Client) ManagementHTTPClient {
	return &ManagementHTTPClientImpl{client}
}

func (c *ManagementHTTPClientImpl) BaseInfo(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*BaseInfoReply, error) {
	var out BaseInfoReply
	pattern := "v1/api/management/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationManagementBaseInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
