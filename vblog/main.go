package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qinchi-ops/govlog/vblog/apps/token/api"
	token "github.com/qinchi-ops/govlog/vblog/apps/token/impl"
	user "github.com/qinchi-ops/govlog/vblog/apps/user/impl"
)

// func main() {
// 	r := gin.Default()
// 	// /api/v1 --> root group
// 	// RouterGroup ---> gin.IRouter
// 	root := r.Group("/vblog/api/v1")

// 	// User BD
// 	userServiceImpl := user.NewUserServiceImpl()
// 	// Token BD
// 	tokenServiceImpl := token.NewTokenServiceImpl(userServiceImpl)
// 	// Token 模块子路由
// 	tokenApiHandler := api.NewTokenApiHandler(tokenServiceImpl)
// 	tokenApiHandler.Registry(root.Group("tokens"))

// 	// fmt.Println("服务在监听 :8080")
// 	if err := r.Run(":8080"); err != nil {
// 		panic(err)
// 	}

// }

func main() {
	//Engine 对HTTP Server的一个封装 Run，也就是整个架构里面协议服务的简化版本
	r := gin.Default()
	// /api/v1 ---> root group
	// RouterGroup ---> gin.IRouter
	root := r.Group("/vblog/api/v1")

	// 把程序需要对象组织起来，组装成业务程序 来处理业务逻辑

	// User BO
	userServiceImpl := user.NewUserServiceImpl()
	// Token BO
	tokenServiceImpl := token.NewTokenServiceImpl(userServiceImpl)

	// token 模块子路有
	tokenApiHandler := api.NewTokenApiHandler(tokenServiceImpl)
	//所有前缀为 /vblog/api/v1/tokens 提交给tokenApiHandler 处理
	tokenApiHandler.Registry(root.Group("tokens"))
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
