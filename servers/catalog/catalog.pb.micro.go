// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: catalog.proto

package catalog

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

// Api Endpoints for Catalog service

func NewCatalogEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Catalog service

type CatalogService interface {
	// 列出所有的产品
	List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*ListRes, error)
	// 增加或者更新产品信息
	Set(ctx context.Context, in *SetReq, opts ...client.CallOption) (*SetRes, error)
	// 删除产品
	Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*DeleteRes, error)
	// 根据产品 ID 查询
	FindByID(ctx context.Context, in *FindByIDReq, opts ...client.CallOption) (*FindByIDRes, error)
}

type catalogService struct {
	c    client.Client
	name string
}

func NewCatalogService(name string, c client.Client) CatalogService {
	return &catalogService{
		c:    c,
		name: name,
	}
}

func (c *catalogService) List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*ListRes, error) {
	req := c.c.NewRequest(c.name, "Catalog.List", in)
	out := new(ListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) Set(ctx context.Context, in *SetReq, opts ...client.CallOption) (*SetRes, error) {
	req := c.c.NewRequest(c.name, "Catalog.Set", in)
	out := new(SetRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*DeleteRes, error) {
	req := c.c.NewRequest(c.name, "Catalog.Delete", in)
	out := new(DeleteRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) FindByID(ctx context.Context, in *FindByIDReq, opts ...client.CallOption) (*FindByIDRes, error) {
	req := c.c.NewRequest(c.name, "Catalog.FindByID", in)
	out := new(FindByIDRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Catalog service

type CatalogHandler interface {
	// 列出所有的产品
	List(context.Context, *ListReq, *ListRes) error
	// 增加或者更新产品信息
	Set(context.Context, *SetReq, *SetRes) error
	// 删除产品
	Delete(context.Context, *DeleteReq, *DeleteRes) error
	// 根据产品 ID 查询
	FindByID(context.Context, *FindByIDReq, *FindByIDRes) error
}

func RegisterCatalogHandler(s server.Server, hdlr CatalogHandler, opts ...server.HandlerOption) error {
	type catalog interface {
		List(ctx context.Context, in *ListReq, out *ListRes) error
		Set(ctx context.Context, in *SetReq, out *SetRes) error
		Delete(ctx context.Context, in *DeleteReq, out *DeleteRes) error
		FindByID(ctx context.Context, in *FindByIDReq, out *FindByIDRes) error
	}
	type Catalog struct {
		catalog
	}
	h := &catalogHandler{hdlr}
	return s.Handle(s.NewHandler(&Catalog{h}, opts...))
}

type catalogHandler struct {
	CatalogHandler
}

func (h *catalogHandler) List(ctx context.Context, in *ListReq, out *ListRes) error {
	return h.CatalogHandler.List(ctx, in, out)
}

func (h *catalogHandler) Set(ctx context.Context, in *SetReq, out *SetRes) error {
	return h.CatalogHandler.Set(ctx, in, out)
}

func (h *catalogHandler) Delete(ctx context.Context, in *DeleteReq, out *DeleteRes) error {
	return h.CatalogHandler.Delete(ctx, in, out)
}

func (h *catalogHandler) FindByID(ctx context.Context, in *FindByIDReq, out *FindByIDRes) error {
	return h.CatalogHandler.FindByID(ctx, in, out)
}
