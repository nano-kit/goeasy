// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: liveuser.proto

package liveuser

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	// 新增或更新用户信息
	AddUser(ctx context.Context, in *AddUserReq, opts ...client.CallOption) (*AddUserRes, error)
	// 查询用户信息
	QueryUser(ctx context.Context, in *QueryUserReq, opts ...client.CallOption) (*QueryUserRes, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) AddUser(ctx context.Context, in *AddUserReq, opts ...client.CallOption) (*AddUserRes, error) {
	req := c.c.NewRequest(c.name, "User.AddUser", in)
	out := new(AddUserRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) QueryUser(ctx context.Context, in *QueryUserReq, opts ...client.CallOption) (*QueryUserRes, error) {
	req := c.c.NewRequest(c.name, "User.QueryUser", in)
	out := new(QueryUserRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 新增或更新用户信息
	AddUser(context.Context, *AddUserReq, *AddUserRes) error
	// 查询用户信息
	QueryUser(context.Context, *QueryUserReq, *QueryUserRes) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		AddUser(ctx context.Context, in *AddUserReq, out *AddUserRes) error
		QueryUser(ctx context.Context, in *QueryUserReq, out *QueryUserRes) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) AddUser(ctx context.Context, in *AddUserReq, out *AddUserRes) error {
	return h.UserHandler.AddUser(ctx, in, out)
}

func (h *userHandler) QueryUser(ctx context.Context, in *QueryUserReq, out *QueryUserRes) error {
	return h.UserHandler.QueryUser(ctx, in, out)
}

// Api Endpoints for Wx service

func NewWxEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Wx service

type WxService interface {
	// 客户端调用 wx.login() 获取临时登录凭证 code ，用此接口回传到开发者服务器。
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginRes, error)
}

type wxService struct {
	c    client.Client
	name string
}

func NewWxService(name string, c client.Client) WxService {
	return &wxService{
		c:    c,
		name: name,
	}
}

func (c *wxService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginRes, error) {
	req := c.c.NewRequest(c.name, "Wx.Login", in)
	out := new(LoginRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Wx service

type WxHandler interface {
	// 客户端调用 wx.login() 获取临时登录凭证 code ，用此接口回传到开发者服务器。
	Login(context.Context, *LoginReq, *LoginRes) error
}

func RegisterWxHandler(s server.Server, hdlr WxHandler, opts ...server.HandlerOption) error {
	type wx interface {
		Login(ctx context.Context, in *LoginReq, out *LoginRes) error
	}
	type Wx struct {
		wx
	}
	h := &wxHandler{hdlr}
	return s.Handle(s.NewHandler(&Wx{h}, opts...))
}

type wxHandler struct {
	WxHandler
}

func (h *wxHandler) Login(ctx context.Context, in *LoginReq, out *LoginRes) error {
	return h.WxHandler.Login(ctx, in, out)
}
