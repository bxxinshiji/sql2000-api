package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/bxxinshiji/sql2000-api/proto/item"
)

// Item 用户结构
type Item struct {
	ServiceName string
}

// Get 获取商品信息
func (srv *Item) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Items.Get", req, res)
}
