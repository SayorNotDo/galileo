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

const OperationCoreCurrentUserInfo = "/api.core.v1.Core/CurrentUserInfo"
const OperationCoreDeleteUser = "/api.core.v1.Core/DeleteUser"
const OperationCoreExecuteToken = "/api.core.v1.Core/ExecuteToken"
const OperationCoreGetUserLatestActivity = "/api.core.v1.Core/GetUserLatestActivity"
const OperationCoreGetUserProjectList = "/api.core.v1.Core/GetUserProjectList"
const OperationCoreListUserGroups = "/api.core.v1.Core/ListUserGroups"
const OperationCoreListUsers = "/api.core.v1.Core/ListUsers"
const OperationCoreLogin = "/api.core.v1.Core/Login"
const OperationCoreLogout = "/api.core.v1.Core/Logout"
const OperationCoreRegister = "/api.core.v1.Core/Register"
const OperationCoreTrackReportData = "/api.core.v1.Core/TrackReportData"
const OperationCoreUpdatePassword = "/api.core.v1.Core/UpdatePassword"
const OperationCoreUserGroups = "/api.core.v1.Core/UserGroups"
const OperationCoreUserInfo = "/api.core.v1.Core/UserInfo"

type CoreHTTPServer interface {
	CurrentUserInfo(context.Context, *emptypb.Empty) (*UserInfoReply, error)
	DeleteUser(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	ExecuteToken(context.Context, *ExecuteTokenRequest) (*ExecuteTokenReply, error)
	GetUserLatestActivity(context.Context, *emptypb.Empty) (*UserLatestActivityReply, error)
	GetUserProjectList(context.Context, *emptypb.Empty) (*UserProjectListReply, error)
	ListUserGroups(context.Context, *ListUserGroupsRequest) (*UserGroupListReply, error)
	ListUsers(context.Context, *ListUserRequest) (*ListUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	TrackReportData(context.Context, *TrackReportDataRequest) (*emptypb.Empty, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error)
	UserGroups(context.Context, *GroupInfoRequest) (*GroupInfo, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
}

func RegisterCoreHTTPServer(s *http.Server, srv CoreHTTPServer) {
	r := s.Route("/")
	r.POST("v1/api/user/register", _Core_Register0_HTTP_Handler(srv))
	r.POST("v1/api/user/login", _Core_Login0_HTTP_Handler(srv))
	r.POST("v1/api/user/logout", _Core_Logout0_HTTP_Handler(srv))
	r.PUT("v1/api/user/{id}/info", _Core_UserInfo0_HTTP_Handler(srv))
	r.GET("v1/api/user/{id}/info", _Core_UserInfo1_HTTP_Handler(srv))
	r.GET("v1/api/user/info", _Core_CurrentUserInfo0_HTTP_Handler(srv))
	r.PUT("v1/api/user/password", _Core_UpdatePassword0_HTTP_Handler(srv))
	r.DELETE("v1/api/user/{id}", _Core_DeleteUser0_HTTP_Handler(srv))
	r.GET("v1/api/user/list/{pageToken}/{pageSize}", _Core_ListUsers0_HTTP_Handler(srv))
	r.PUT("v1/api/user/groups/{id}", _Core_UserGroups0_HTTP_Handler(srv))
	r.GET("v1/api/user/groups/{id}", _Core_UserGroups1_HTTP_Handler(srv))
	r.DELETE("v1/api/user/groups/{id}", _Core_UserGroups2_HTTP_Handler(srv))
	r.POST("v1/api/user/group", _Core_UserGroups3_HTTP_Handler(srv))
	r.GET("v1/api/user/groups", _Core_ListUserGroups0_HTTP_Handler(srv))
	r.GET("v1/api/user/project/list", _Core_GetUserProjectList0_HTTP_Handler(srv))
	r.GET("v1/api/user/latest-activity", _Core_GetUserLatestActivity0_HTTP_Handler(srv))
	r.POST("data/report", _Core_TrackReportData0_HTTP_Handler(srv))
	r.POST("v1/api/execute-token", _Core_ExecuteToken0_HTTP_Handler(srv))
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

func _Core_UserInfo0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
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
		if err := ctx.BindVars(&in); err != nil {
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

func _Core_CurrentUserInfo0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreCurrentUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CurrentUserInfo(ctx, req.(*emptypb.Empty))
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
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Core_ListUsers0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreListUsers)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUsers(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroups0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
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
		http.SetOperation(ctx, OperationCoreUserGroups)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroups(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroups1_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserGroups)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroups(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroups2_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserGroups)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroups(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_UserGroups3_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserGroups)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserGroups(ctx, req.(*GroupInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupInfo)
		return ctx.Result(200, reply)
	}
}

func _Core_ListUserGroups0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserGroupsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreListUserGroups)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserGroups(ctx, req.(*ListUserGroupsRequest))
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

func _Core_TrackReportData0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TrackReportDataRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreTrackReportData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TrackReportData(ctx, req.(*TrackReportDataRequest))
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

type CoreHTTPClient interface {
	CurrentUserInfo(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	DeleteUser(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ExecuteToken(ctx context.Context, req *ExecuteTokenRequest, opts ...http.CallOption) (rsp *ExecuteTokenReply, err error)
	GetUserLatestActivity(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserLatestActivityReply, err error)
	GetUserProjectList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserProjectListReply, err error)
	ListUserGroups(ctx context.Context, req *ListUserGroupsRequest, opts ...http.CallOption) (rsp *UserGroupListReply, err error)
	ListUsers(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	TrackReportData(ctx context.Context, req *TrackReportDataRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdatePassword(ctx context.Context, req *UpdatePasswordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UserGroups(ctx context.Context, req *GroupInfoRequest, opts ...http.CallOption) (rsp *GroupInfo, err error)
	UserInfo(ctx context.Context, req *UserInfoRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
}

type CoreHTTPClientImpl struct {
	cc *http.Client
}

func NewCoreHTTPClient(client *http.Client) CoreHTTPClient {
	return &CoreHTTPClientImpl{client}
}

func (c *CoreHTTPClientImpl) CurrentUserInfo(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "v1/api/user/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreCurrentUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
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

func (c *CoreHTTPClientImpl) ListUserGroups(ctx context.Context, in *ListUserGroupsRequest, opts ...http.CallOption) (*UserGroupListReply, error) {
	var out UserGroupListReply
	pattern := "v1/api/user/groups"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreListUserGroups))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) ListUsers(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "v1/api/user/list/{pageToken}/{pageSize}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreListUsers))
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

func (c *CoreHTTPClientImpl) TrackReportData(ctx context.Context, in *TrackReportDataRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "data/report"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreTrackReportData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
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

func (c *CoreHTTPClientImpl) UserGroups(ctx context.Context, in *GroupInfoRequest, opts ...http.CallOption) (*GroupInfo, error) {
	var out GroupInfo
	pattern := "v1/api/user/group"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUserGroups))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "v1/api/user/{id}/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
