// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
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

const OperationCoreCreateGroup = "/api.core.v1.Core/CreateGroup"
const OperationCoreDataReportTrack = "/api.core.v1.Core/DataReportTrack"
const OperationCoreDeleteUser = "/api.core.v1.Core/DeleteUser"
const OperationCoreExecuteToken = "/api.core.v1.Core/ExecuteToken"
const OperationCoreInspectContainer = "/api.core.v1.Core/InspectContainer"
const OperationCoreListUser = "/api.core.v1.Core/ListUser"
const OperationCoreLogin = "/api.core.v1.Core/Login"
const OperationCoreLogout = "/api.core.v1.Core/Logout"
const OperationCoreRegister = "/api.core.v1.Core/Register"
const OperationCoreUnregister = "/api.core.v1.Core/Unregister"
const OperationCoreUpdateEmail = "/api.core.v1.Core/UpdateEmail"
const OperationCoreUpdateGroup = "/api.core.v1.Core/UpdateGroup"
const OperationCoreUpdatePassword = "/api.core.v1.Core/UpdatePassword"
const OperationCoreUpdatePhone = "/api.core.v1.Core/UpdatePhone"
const OperationCoreUpdateUserInfo = "/api.core.v1.Core/UpdateUserInfo"
const OperationCoreUserDetail = "/api.core.v1.Core/UserDetail"

type CoreHTTPServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupReply, error)
	DataReportTrack(context.Context, *DataReportTrackRequest) (*emptypb.Empty, error)
	DeleteUser(context.Context, *DeleteRequest) (*DeleteReply, error)
	ExecuteToken(context.Context, *ExecuteTokenRequest) (*ExecuteTokenReply, error)
	InspectContainer(context.Context, *InspectContainerRequest) (*InspectContainerReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Unregister(context.Context, *UnregisterRequest) (*UnregisterReply, error)
	UpdateEmail(context.Context, *UpdateEmailRequest) (*UpdateEmailReply, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*emptypb.Empty, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error)
	UpdatePhone(context.Context, *UpdatePhoneRequest) (*UpdatePhoneReply, error)
	UpdateUserInfo(context.Context, *UserInfoUpdateRequest) (*UserInfoUpdateReply, error)
	UserDetail(context.Context, *emptypb.Empty) (*UserDetailReply, error)
}

func RegisterCoreHTTPServer(s *http.Server, srv CoreHTTPServer) {
	r := s.Route("/")
	r.POST("v1/api/user/register", _Core_Register0_HTTP_Handler(srv))
	r.POST("v1/api/user/login", _Core_Login0_HTTP_Handler(srv))
	r.DELETE("v1/api/user/unregister", _Core_Unregister0_HTTP_Handler(srv))
	r.POST("v1/api/user/logout", _Core_Logout0_HTTP_Handler(srv))
	r.DELETE("v1/api/user/delete/{id}", _Core_DeleteUser0_HTTP_Handler(srv))
	r.GET("v1/api/user/info", _Core_UserDetail0_HTTP_Handler(srv))
	r.PUT("v1/api/user/update", _Core_UpdateUserInfo0_HTTP_Handler(srv))
	r.PUT("v1/api/user/password", _Core_UpdatePassword0_HTTP_Handler(srv))
	r.PUT("v1/api/user/email", _Core_UpdateEmail0_HTTP_Handler(srv))
	r.PUT("v1/api/user/phone", _Core_UpdatePhone0_HTTP_Handler(srv))
	r.GET("v1/api/user/list/{pageNum}/{pageSize}", _Core_ListUser0_HTTP_Handler(srv))
	r.POST("v1/api/user/group", _Core_CreateGroup0_HTTP_Handler(srv))
	r.PUT("v1/api/user/group", _Core_UpdateGroup0_HTTP_Handler(srv))
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
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
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

func _Core_UserDetail0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUserDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserDetail(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserDetailReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UpdateUserInfo0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUpdateUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserInfo(ctx, req.(*UserInfoUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UpdatePassword0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePasswordRequest
		if err := ctx.Bind(&in); err != nil {
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

func _Core_UpdateEmail0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateEmailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUpdateEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateEmail(ctx, req.(*UpdateEmailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateEmailReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UpdatePhone0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePhoneRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUpdatePhone)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePhone(ctx, req.(*UpdatePhoneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePhoneReply)
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

func _Core_CreateGroup0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreCreateGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateGroup(ctx, req.(*CreateGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Core_UpdateGroup0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCoreUpdateGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateGroup(ctx, req.(*UpdateGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Core_DataReportTrack0_HTTP_Handler(srv CoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DataReportTrackRequest
		if err := ctx.Bind(&in); err != nil {
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
		reply := out.(*InspectContainerReply)
		return ctx.Result(200, reply)
	}
}

type CoreHTTPClient interface {
	CreateGroup(ctx context.Context, req *CreateGroupRequest, opts ...http.CallOption) (rsp *CreateGroupReply, err error)
	DataReportTrack(ctx context.Context, req *DataReportTrackRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteUser(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *DeleteReply, err error)
	ExecuteToken(ctx context.Context, req *ExecuteTokenRequest, opts ...http.CallOption) (rsp *ExecuteTokenReply, err error)
	InspectContainer(ctx context.Context, req *InspectContainerRequest, opts ...http.CallOption) (rsp *InspectContainerReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	Unregister(ctx context.Context, req *UnregisterRequest, opts ...http.CallOption) (rsp *UnregisterReply, err error)
	UpdateEmail(ctx context.Context, req *UpdateEmailRequest, opts ...http.CallOption) (rsp *UpdateEmailReply, err error)
	UpdateGroup(ctx context.Context, req *UpdateGroupRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdatePassword(ctx context.Context, req *UpdatePasswordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdatePhone(ctx context.Context, req *UpdatePhoneRequest, opts ...http.CallOption) (rsp *UpdatePhoneReply, err error)
	UpdateUserInfo(ctx context.Context, req *UserInfoUpdateRequest, opts ...http.CallOption) (rsp *UserInfoUpdateReply, err error)
	UserDetail(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserDetailReply, err error)
}

type CoreHTTPClientImpl struct {
	cc *http.Client
}

func NewCoreHTTPClient(client *http.Client) CoreHTTPClient {
	return &CoreHTTPClientImpl{client}
}

func (c *CoreHTTPClientImpl) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...http.CallOption) (*CreateGroupReply, error) {
	var out CreateGroupReply
	pattern := "v1/api/user/group"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreCreateGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
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
	pattern := "v1/api/user/delete/{id}"
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

func (c *CoreHTTPClientImpl) InspectContainer(ctx context.Context, in *InspectContainerRequest, opts ...http.CallOption) (*InspectContainerReply, error) {
	var out InspectContainerReply
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

func (c *CoreHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
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

func (c *CoreHTTPClientImpl) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...http.CallOption) (*UpdateEmailReply, error) {
	var out UpdateEmailReply
	pattern := "v1/api/user/email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUpdateEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "v1/api/user/group"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUpdateGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
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

func (c *CoreHTTPClientImpl) UpdatePhone(ctx context.Context, in *UpdatePhoneRequest, opts ...http.CallOption) (*UpdatePhoneReply, error) {
	var out UpdatePhoneReply
	pattern := "v1/api/user/phone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUpdatePhone))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) UpdateUserInfo(ctx context.Context, in *UserInfoUpdateRequest, opts ...http.CallOption) (*UserInfoUpdateReply, error) {
	var out UserInfoUpdateReply
	pattern := "v1/api/user/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCoreUpdateUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoreHTTPClientImpl) UserDetail(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserDetailReply, error) {
	var out UserDetailReply
	pattern := "v1/api/user/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCoreUserDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
