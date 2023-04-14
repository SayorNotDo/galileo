// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.7
// source: api/management/testcase/v1/testcase.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTestcaseCreateTestcase = "/api.management.testcase.Testcase/CreateTestcase"

type TestcaseHTTPServer interface {
	CreateTestcase(context.Context, *CreateTestcaseRequest) (*CreateTestcaseReply, error)
}

func RegisterTestcaseHTTPServer(s *http.Server, srv TestcaseHTTPServer) {
	r := s.Route("/")
	r.POST("v1/api/testcase", _Testcase_CreateTestcase0_HTTP_Handler(srv))
}

func _Testcase_CreateTestcase0_HTTP_Handler(srv TestcaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTestcaseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTestcaseCreateTestcase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTestcase(ctx, req.(*CreateTestcaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTestcaseReply)
		return ctx.Result(200, reply)
	}
}

type TestcaseHTTPClient interface {
	CreateTestcase(ctx context.Context, req *CreateTestcaseRequest, opts ...http.CallOption) (rsp *CreateTestcaseReply, err error)
}

type TestcaseHTTPClientImpl struct {
	cc *http.Client
}

func NewTestcaseHTTPClient(client *http.Client) TestcaseHTTPClient {
	return &TestcaseHTTPClientImpl{client}
}

func (c *TestcaseHTTPClientImpl) CreateTestcase(ctx context.Context, in *CreateTestcaseRequest, opts ...http.CallOption) (*CreateTestcaseReply, error) {
	var out CreateTestcaseReply
	pattern := "v1/api/testcase"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTestcaseCreateTestcase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
