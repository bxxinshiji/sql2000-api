package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/util/log"

	server "github.com/micro/go-micro/v2/server"

	departmentPB "github.com/bxxinshiji/sql2000-api/proto/department"
	healthPB "github.com/bxxinshiji/sql2000-api/proto/health"
	itemPB "github.com/bxxinshiji/sql2000-api/proto/item"

	client "github.com/lecex/core/client"

	"github.com/bxxinshiji/sql2000-api/config"
	PB "github.com/lecex/user/proto/permission"
)

var Conf = config.Conf

// Register 注册
func Register(Server server.Server) {
	itemPB.RegisterItemsHandler(Server, &Item{Conf.Service["sql2000"]}) // 权限管理服务实现
	departmentPB.RegisterDepartmentHandler(Server, &Department{Conf.Service["sql2000"]})
	healthPB.RegisterHealthHandler(Server, &Health{})

	go Sync() // 同步前端权限
}

// Sync 同步
func Sync() {
	time.Sleep(5 * time.Second)
	req := &PB.Request{
		Permissions: Conf.Permissions,
	}
	res := &PB.Response{}
	err := client.Call(context.TODO(), Conf.Service["user"], "Permissions.Sync", req, res)
	if err != nil {
		log.Log(err)
		Sync()
	}
}
