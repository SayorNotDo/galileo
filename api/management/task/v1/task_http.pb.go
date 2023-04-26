// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.7
// source: api/management/task/v1/task.proto

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

const OperationTaskCreateTask = "/api.task.v1.Task/CreateTask"
const OperationTaskTaskByID = "/api.task.v1.Task/TaskByID"
const OperationTaskTestEngine = "/api.task.v1.Task/TestEngine"
const OperationTaskUpdateTask = "/api.task.v1.Task/UpdateTask"

type TaskHTTPServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskReply, error)
	TaskByID(context.Context, *TaskByIDRequest) (*GetTaskReply, error)
	TestEngine(context.Context, *emptypb.Empty) (*TestEngineReply, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskReply, error)
}

func RegisterTaskHTTPServer(s *http.Server, srv TaskHTTPServer) {
	r := s.Route("/")
	r.POST("v1/api/management/task", _Task_CreateTask0_HTTP_Handler(srv))
	r.PUT("v1/api/management/task", _Task_UpdateTask0_HTTP_Handler(srv))
	r.GET("v1/api/management/task/{id}", _Task_TaskByID0_HTTP_Handler(srv))
	r.GET("v1/api/engine/ciao", _Task_TestEngine0_HTTP_Handler(srv))
}

func _Task_CreateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTask(ctx, req.(*CreateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_TaskByID0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskByIDRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskTaskByID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TaskByID(ctx, req.(*TaskByIDRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_TestEngine0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskTestEngine)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TestEngine(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestEngineReply)
		return ctx.Result(200, reply)
	}
}

type TaskHTTPClient interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest, opts ...http.CallOption) (rsp *CreateTaskReply, err error)
	TaskByID(ctx context.Context, req *TaskByIDRequest, opts ...http.CallOption) (rsp *GetTaskReply, err error)
	TestEngine(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *TestEngineReply, err error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest, opts ...http.CallOption) (rsp *UpdateTaskReply, err error)
}

type TaskHTTPClientImpl struct {
	cc *http.Client
}

func NewTaskHTTPClient(client *http.Client) TaskHTTPClient {
	return &TaskHTTPClientImpl{client}
}

func (c *TaskHTTPClientImpl) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...http.CallOption) (*CreateTaskReply, error) {
	var out CreateTaskReply
	pattern := "v1/api/management/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) TaskByID(ctx context.Context, in *TaskByIDRequest, opts ...http.CallOption) (*GetTaskReply, error) {
	var out GetTaskReply
	pattern := "v1/api/management/task/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskTaskByID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) TestEngine(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*TestEngineReply, error) {
	var out TestEngineReply
	pattern := "v1/api/engine/ciao"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskTestEngine))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...http.CallOption) (*UpdateTaskReply, error) {
	var out UpdateTaskReply
	pattern := "v1/api/management/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
