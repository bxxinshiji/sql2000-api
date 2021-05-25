// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/item/item.proto

package item

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Items service

type ItemsService interface {
	// 查询商品详情
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type itemsService struct {
	c    client.Client
	name string
}

func NewItemsService(name string, c client.Client) ItemsService {
	return &itemsService{
		c:    c,
		name: name,
	}
}

func (c *itemsService) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Items.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsService) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Items.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Items service

type ItemsHandler interface {
	// 查询商品详情
	All(context.Context, *Request, *Response) error
	Get(context.Context, *Request, *Response) error
}

func RegisterItemsHandler(s server.Server, hdlr ItemsHandler, opts ...server.HandlerOption) error {
	type items interface {
		All(ctx context.Context, in *Request, out *Response) error
		Get(ctx context.Context, in *Request, out *Response) error
	}
	type Items struct {
		items
	}
	h := &itemsHandler{hdlr}
	return s.Handle(s.NewHandler(&Items{h}, opts...))
}

type itemsHandler struct {
	ItemsHandler
}

func (h *itemsHandler) All(ctx context.Context, in *Request, out *Response) error {
	return h.ItemsHandler.All(ctx, in, out)
}

func (h *itemsHandler) Get(ctx context.Context, in *Request, out *Response) error {
	return h.ItemsHandler.Get(ctx, in, out)
}
