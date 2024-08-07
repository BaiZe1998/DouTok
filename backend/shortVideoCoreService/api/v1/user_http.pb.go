// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: v1/user.proto

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

const OperationShortVideoCoreUserServiceGetUserInfo = "/shortVideoCoreService.api.v1.ShortVideoCoreUserService/GetUserInfo"
const OperationShortVideoCoreUserServiceLogin = "/shortVideoCoreService.api.v1.ShortVideoCoreUserService/Login"
const OperationShortVideoCoreUserServiceRegister = "/shortVideoCoreService.api.v1.ShortVideoCoreUserService/Register"
const OperationShortVideoCoreUserServiceUpdateUserInfo = "/shortVideoCoreService.api.v1.ShortVideoCoreUserService/UpdateUserInfo"

type ShortVideoCoreUserServiceHTTPServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
}

func RegisterShortVideoCoreUserServiceHTTPServer(s *http.Server, srv ShortVideoCoreUserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user/register", _ShortVideoCoreUserService_Register0_HTTP_Handler(srv))
	r.POST("/v1/user/login", _ShortVideoCoreUserService_Login0_HTTP_Handler(srv))
	r.GET("/v1/user/info", _ShortVideoCoreUserService_GetUserInfo0_HTTP_Handler(srv))
	r.PUT("/v1/user/info", _ShortVideoCoreUserService_UpdateUserInfo0_HTTP_Handler(srv))
}

func _ShortVideoCoreUserService_Register0_HTTP_Handler(srv ShortVideoCoreUserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortVideoCoreUserServiceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResponse)
		return ctx.Result(200, reply)
	}
}

func _ShortVideoCoreUserService_Login0_HTTP_Handler(srv ShortVideoCoreUserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortVideoCoreUserServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _ShortVideoCoreUserService_GetUserInfo0_HTTP_Handler(srv ShortVideoCoreUserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortVideoCoreUserServiceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ShortVideoCoreUserService_UpdateUserInfo0_HTTP_Handler(srv ShortVideoCoreUserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortVideoCoreUserServiceUpdateUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

type ShortVideoCoreUserServiceHTTPClient interface {
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterResponse, err error)
	UpdateUserInfo(ctx context.Context, req *UpdateUserInfoRequest, opts ...http.CallOption) (rsp *UpdateUserInfoResponse, err error)
}

type ShortVideoCoreUserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewShortVideoCoreUserServiceHTTPClient(client *http.Client) ShortVideoCoreUserServiceHTTPClient {
	return &ShortVideoCoreUserServiceHTTPClientImpl{client}
}

func (c *ShortVideoCoreUserServiceHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoResponse, error) {
	var out GetUserInfoResponse
	pattern := "/v1/user/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShortVideoCoreUserServiceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ShortVideoCoreUserServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShortVideoCoreUserServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ShortVideoCoreUserServiceHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterResponse, error) {
	var out RegisterResponse
	pattern := "/v1/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShortVideoCoreUserServiceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ShortVideoCoreUserServiceHTTPClientImpl) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...http.CallOption) (*UpdateUserInfoResponse, error) {
	var out UpdateUserInfoResponse
	pattern := "/v1/user/info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShortVideoCoreUserServiceUpdateUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
