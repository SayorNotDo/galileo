// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.21.12
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
const OperationTestcaseCreateTestcaseSuite = "/api.management.testcase.Testcase/CreateTestcaseSuite"
const OperationTestcaseDebugTestcase = "/api.management.testcase.Testcase/DebugTestcase"
const OperationTestcaseGetTestcaseById = "/api.management.testcase.Testcase/GetTestcaseById"
const OperationTestcaseLoadFramework = "/api.management.testcase.Testcase/LoadFramework"
const OperationTestcaseUpdateTestcase = "/api.management.testcase.Testcase/UpdateTestcase"

type TestcaseHTTPServer interface {
	CreateTestcase(context.Context, *CreateTestcaseRequest) (*CreateTestcaseReply, error)
	CreateTestcaseSuite(context.Context, *CreateTestcaseSuiteRequest) (*CreateTestcaseSuiteReply, error)
	DebugTestcase(context.Context, *DebugTestcaseRequest) (*DebugTestcaseReply, error)
	GetTestcaseById(context.Context, *GetTestcaseRequest) (*GetTestcaseReply, error)
	LoadFramework(context.Context, *LoadFrameworkRequest) (*LoadFrameworkReply, error)
	UpdateTestcase(context.Context, *UpdateTestcaseRequest) (*UpdateTestcaseReply, error)
}

func RegisterTestcaseHTTPServer(s *http.Server, srv TestcaseHTTPServer) {
	r := s.Route("/")
	r.POST("v1/api/management/testcase", _Testcase_CreateTestcase0_HTTP_Handler(srv))
	r.PUT("v1/api/management/testcase", _Testcase_UpdateTestcase0_HTTP_Handler(srv))
	r.GET("v1/api/management/testcase/{id}", _Testcase_GetTestcaseById0_HTTP_Handler(srv))
	r.POST("v1/api/management/testcase/debug", _Testcase_DebugTestcase0_HTTP_Handler(srv))
	r.POST("v1/api/management/testcase/loadFramework", _Testcase_LoadFramework0_HTTP_Handler(srv))
	r.POST("v1/api/management/case-suite", _Testcase_CreateTestcaseSuite0_HTTP_Handler(srv))
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

func _Testcase_UpdateTestcase0_HTTP_Handler(srv TestcaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTestcaseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTestcaseUpdateTestcase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTestcase(ctx, req.(*UpdateTestcaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTestcaseReply)
		return ctx.Result(200, reply)
	}
}

func _Testcase_GetTestcaseById0_HTTP_Handler(srv TestcaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTestcaseRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTestcaseGetTestcaseById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTestcaseById(ctx, req.(*GetTestcaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTestcaseReply)
		return ctx.Result(200, reply)
	}
}

func _Testcase_DebugTestcase0_HTTP_Handler(srv TestcaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DebugTestcaseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTestcaseDebugTestcase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DebugTestcase(ctx, req.(*DebugTestcaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DebugTestcaseReply)
		return ctx.Result(200, reply)
	}
}

func _Testcase_LoadFramework0_HTTP_Handler(srv TestcaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoadFrameworkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTestcaseLoadFramework)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoadFramework(ctx, req.(*LoadFrameworkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoadFrameworkReply)
		return ctx.Result(200, reply)
	}
}

func _Testcase_CreateTestcaseSuite0_HTTP_Handler(srv TestcaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTestcaseSuiteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTestcaseCreateTestcaseSuite)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTestcaseSuite(ctx, req.(*CreateTestcaseSuiteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTestcaseSuiteReply)
		return ctx.Result(200, reply)
	}
}

type TestcaseHTTPClient interface {
	CreateTestcase(ctx context.Context, req *CreateTestcaseRequest, opts ...http.CallOption) (rsp *CreateTestcaseReply, err error)
	CreateTestcaseSuite(ctx context.Context, req *CreateTestcaseSuiteRequest, opts ...http.CallOption) (rsp *CreateTestcaseSuiteReply, err error)
	DebugTestcase(ctx context.Context, req *DebugTestcaseRequest, opts ...http.CallOption) (rsp *DebugTestcaseReply, err error)
	GetTestcaseById(ctx context.Context, req *GetTestcaseRequest, opts ...http.CallOption) (rsp *GetTestcaseReply, err error)
	LoadFramework(ctx context.Context, req *LoadFrameworkRequest, opts ...http.CallOption) (rsp *LoadFrameworkReply, err error)
	UpdateTestcase(ctx context.Context, req *UpdateTestcaseRequest, opts ...http.CallOption) (rsp *UpdateTestcaseReply, err error)
}

type TestcaseHTTPClientImpl struct {
	cc *http.Client
}

func NewTestcaseHTTPClient(client *http.Client) TestcaseHTTPClient {
	return &TestcaseHTTPClientImpl{client}
}

func (c *TestcaseHTTPClientImpl) CreateTestcase(ctx context.Context, in *CreateTestcaseRequest, opts ...http.CallOption) (*CreateTestcaseReply, error) {
	var out CreateTestcaseReply
	pattern := "v1/api/management/testcase"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTestcaseCreateTestcase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TestcaseHTTPClientImpl) CreateTestcaseSuite(ctx context.Context, in *CreateTestcaseSuiteRequest, opts ...http.CallOption) (*CreateTestcaseSuiteReply, error) {
	var out CreateTestcaseSuiteReply
	pattern := "v1/api/management/case-suite"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTestcaseCreateTestcaseSuite))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TestcaseHTTPClientImpl) DebugTestcase(ctx context.Context, in *DebugTestcaseRequest, opts ...http.CallOption) (*DebugTestcaseReply, error) {
	var out DebugTestcaseReply
	pattern := "v1/api/management/testcase/debug"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTestcaseDebugTestcase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TestcaseHTTPClientImpl) GetTestcaseById(ctx context.Context, in *GetTestcaseRequest, opts ...http.CallOption) (*GetTestcaseReply, error) {
	var out GetTestcaseReply
	pattern := "v1/api/management/testcase/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTestcaseGetTestcaseById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TestcaseHTTPClientImpl) LoadFramework(ctx context.Context, in *LoadFrameworkRequest, opts ...http.CallOption) (*LoadFrameworkReply, error) {
	var out LoadFrameworkReply
	pattern := "v1/api/management/testcase/loadFramework"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTestcaseLoadFramework))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TestcaseHTTPClientImpl) UpdateTestcase(ctx context.Context, in *UpdateTestcaseRequest, opts ...http.CallOption) (*UpdateTestcaseReply, error) {
	var out UpdateTestcaseReply
	pattern := "v1/api/management/testcase"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTestcaseUpdateTestcase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
