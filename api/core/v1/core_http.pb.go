// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.12
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

const OperationCoreDataReportTrack = "/api.core.v1.Core/DataReportTrack"
const OperationCoreDeleteUser = "/api.core.v1.Core/DeleteUser"
const OperationCoreExecuteToken = "/api.core.v1.Core/ExecuteToken"
const OperationCoreGetUserGroupList = "/api.core.v1.Core/GetUserGroupList"
const OperationCoreGetUserLatestActivity = "/api.core.v1.Core/GetUserLatestActivity"
const OperationCoreGetUserProjectList = "/api.core.v1.Core/GetUserProjectList"
const OperationCoreInspectContainer = "/api.core.v1.Core/InspectContainer"
const OperationCoreListUser = "/api.core.v1.Core/ListUser"
const OperationCoreLogin = "/api.core.v1.Core/Login"
const OperationCoreLogout = "/api.core.v1.Core/Logout"
const OperationCoreRegister = "/api.core.v1.Core/Register"
const OperationCoreUnregister = "/api.core.v1.Core/Unregister"
const OperationCoreUpdatePassword = "/api.core.v1.Core/UpdatePassword"
const OperationCoreUserGroup = "/api.core.v1.Core/UserGroup"
const OperationCoreUserInfo = "/api.core.v1.Core/UserInfo"

type CoreHTTPServer interface {
	DataReportTrack(context.Context, *DataReportTrackRequest) (*emptypb.Empty, error)
	DeleteUser(context.Context, *DeleteRequest) (*DeleteReply, error)
	ExecuteToken(context.Context, *ExecuteTokenRequest) (*ExecuteTokenReply, error)
	GetUserGroupList(context.Context, *emptypb.Empty) (*UserGroupListReply, error)
	GetUserLatestActivity(context.Context, *emptypb.Empty) (*UserLatestActivityReply, error)
	GetUserProjectList(context.Context, *emptypb.Empty) (*UserProjectListReply, error)
	InspectContainer(context.Context, *InspectContainerRequest) (*ContainerInfo, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Unregister(context.Context, *UnregisterRequest) (*UnregisterReply, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error)
	UserGroup(context.Context, *GroupInfoRequest) (*GroupInfo, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
}

func RegisterCoreHTTPServer(s *http.Server, srv CoreHTTPServer) {
	r := s.Route("/")
	r.POST("v1/api/user/register", _Core_Register0_HTTP_Handler(srv))
	r.POST("v1/api/user/login", _Core_Login0_HTTP_Handler(srv))
	r.DELETE("v1/api/user/unregister", _Core_Unregister0_HTTP_Handler(srv))
	r.POST("v1/api/user/logout", _Core_Logout0_HTTP_Handler(srv))
	r.DELETE("v1/api/user/{id}", _Core_DeleteUser0_HTTP_Handler(srv))
	r.PUT("v1/api/user/info", _Core_UserInfo0_HTTP_Handler(srv))
	r.GET("v1/api/user/info", _Core_UserInfo1_HTTP_Handler(srv))
	r.PUT("v1/api/user/password", _Core_UpdatePassword0_HTTP_Handler(srv))
	r.GET("v1/api/user/list/{pageNum}/{pageSize}", _Core_ListUser0_HTTP_Handler(srv))
	r.PUT("v1/api/user/group/{id}", _Core_UserGroup0_HTTP_Handler(srv))
	r.GET("v1/api/user/group/{id}", _Core_UserGroup1_HTTP_Handler(srv))
	r.DELETE("v1/api/user/group/{id}", _Core_UserGroup2_HTTP_Handler(srv))
	r.POST("v1/api/user/group", _Core_UserGroup3_HTTP_Handler(srv))
	r.GET("v1/api/user/group/list", _Core_GetUserGroupList0_HTTP_Handler(srv))
	r.GET("v1/api/user/project/list", _Core_GetUserProjectList0_HTTP_Handler(srv))
	r.GET("v1/api/user/latest-activity", _Core_GetUserLatestActivity0_HTTP_Handler(srv))
	r.POST("data/report", _Core_DataReportTrack0_HTTP_Handler(srv))
	r.POST("v1/api/execute-token", _Core_ExecuteToken0_HTTP_Handler(srv))
	r.GET("v1/api/engine/container/{id}", _Core_InspectContainer0_HTTP_Handler(srv))
}

func _Core_Register0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
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
		if err := ctx.BindQuery(&in); err != nil {
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
		if err := ctx.BindQuery(&in); err != nil {
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
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Core_DeleteUser0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UserInfo0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfo(ctx, req.(*UserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UserInfo1_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfo(ctx, req.(*UserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UpdatePassword0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUpdatePassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePassword(ctx, req.(*UpdatePasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Core_ListUser0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroup0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroup(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroup1_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroup(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroup2_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroup(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroup3_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroup(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_GetUserGroupList0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreGetUserGroupList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserGroupList(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserGroupListReply)
		return ctx.Result(200, reply)
	}
}

func _Core_GetUserProjectList0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreGetUserProjectList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserProjectList(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserProjectListReply)
		return ctx.Result(200, reply)
	}
}

func _Core_GetUserLatestActivity0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreGetUserLatestActivity)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserLatestActivity(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserLatestActivityReply)
		return ctx.Result(200, reply)
	}
}

func _Core_DataReportTrack0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DataReportTrackRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreDataReportTrack)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DataReportTrack(ctx, req.(*DataReportTrackRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Core_ExecuteToken0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExecuteTokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreExecuteToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExecuteToken(ctx, req.(*ExecuteTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExecuteTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Core_InspectContainer0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InspectContainerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreInspectContainer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InspectContainer(ctx, req.(*InspectContainerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ContainerInfo)
		return ctx.Result(200, reply)
	}
}

type CoreHTTPClient interface {
	DataReportTrack(ctx context.Context, req *DataReportTrackRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteUser(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *DeleteReply, err error)
	ExecuteToken(ctx context.Context, req *ExecuteTokenRequest, opts ...http.CallOption) (rsp *ExecuteTokenReply, err error)
	GetUserGroupList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserGroupListReply, err error)
	GetUserLatestActivity(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserLatestActivityReply, err error)
	GetUserProjectList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserProjectListReply, err error)
	InspectContainer(ctx context.Context, req *InspectContainerRequest, opts ...http.CallOption) (rsp *ContainerInfo, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	Unregister(ctx context.Context, req *UnregisterRequest, opts ...http.CallOption) (rsp *UnregisterReply, err error)
	UpdatePassword(ctx context.Context, req *UpdatePasswordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UserGroup(ctx context.Context, req *GroupInfoRequest, opts ...http.CallOption) (rsp *GroupInfo, err error)
	UserInfo(ctx context.Context, req *UserInfoRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
}

type CoreHTTPClientImpl struct {
	cc *http.Client
}

func NewCoreHTTPClient(client *http.Client) CoreHTTPClient {
	return &CoreHTTPClientImpl{client}
}

func (c *CoreHTTPClientImpl) DataReportTrack(ctx context.Context, in *DataReportTrackRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "data/report"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreDataReportTrack))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*DeleteReply, error) {
	var out DeleteReply
	pattern := "v1/api/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) ExecuteToken(ctx context.Context, in *ExecuteTokenRequest, opts ...http.CallOption) (*ExecuteTokenReply, error) {
	var out ExecuteTokenReply
	pattern := "v1/api/execute-token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreExecuteToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) GetUserGroupList(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserGroupListReply, error) {
	var out UserGroupListReply
	pattern := "v1/api/user/group/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreGetUserGroupList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) GetUserLatestActivity(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserLatestActivityReply, error) {
	var out UserLatestActivityReply
	pattern := "v1/api/user/latest-activity"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreGetUserLatestActivity))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) GetUserProjectList(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserProjectListReply, error) {
	var out UserProjectListReply
	pattern := "v1/api/user/project/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreGetUserProjectList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) InspectContainer(ctx context.Context, in *InspectContainerRequest, opts ...http.CallOption) (*ContainerInfo, error) {
	var out ContainerInfo
	pattern := "v1/api/engine/container/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreInspectContainer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "v1/api/user/list/{pageNum}/{pageSize}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreListUser))
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

func (c *CoreHTTPClientImpl) Logout(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
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

func (c *CoreHTTPClientImpl) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "v1/api/user/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUpdatePassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) UserGroup(ctx context.Context, in *GroupInfoRequest, opts ...http.CallOption) (*GroupInfo, error) {
	var out GroupInfo
	pattern := "v1/api/user/group"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUserGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "v1/api/user/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
