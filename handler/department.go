package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/bxxinshiji/sql2000-api/proto/department"
)

// Department 用户结构
type Department struct {
	ServiceName string
}

// Sale 获取商品信息
func (srv *Department) Sale(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Department.Sale", req, res)
}
