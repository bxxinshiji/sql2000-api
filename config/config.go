package config

import (
	"github.com/lecex/core/config"
	"github.com/lecex/core/env"
	PB "github.com/lecex/user/proto/permission"
)

// 	Conf 配置
// 	Name // 服务名称
//	Method // 方法
//	Auth // 是否认证授权
//	Policy // 是否认证权限
//	Name // 权限名称
//	Description // 权限解释
var Conf config.Config = config.Config{
	Name:    env.Getenv("MICRO_API_NAMESPACE", "go.micro.api.") + "sql2000-api",
	Version: "latest",
	Service: map[string]string{
		"user":    env.Getenv("USER_SERVICE", "go.micro.srv.user"),
		"sql2000": env.Getenv("SQL2000_SERVICE", "go.micro.srv.sql2000"),
	},
	Permissions: []*PB.Permission{
		// 商品
		{Service: "sql2000-api", Method: "Items.Get", Auth: true, Policy: false, Name: "获取商品详细", Description: "获取商品详细,包括多条码信息。"},
		{Service: "sql2000-api", Method: "Department.Sale", Auth: true, Policy: true, Name: "部门销售额", Description: "部门销售额查询。"},
	},
}
