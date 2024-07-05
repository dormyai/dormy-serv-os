// main.go
package main

import (
	adminRouter "dormy/app/admin/router"
	frontendRouter "dormy/app/frontend/router"
	"dormy/app/listening"
	"dormy/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gin.SetMode(gin.ReleaseMode)

	// 加载前台用户模块路由
	frontendRouter.Load(r.Group("/api/v1"))

	// 加载后台用户模块路由
	adminRouter.Load(r.Group("/api/admin/v1"))

	go listening.ListenToDormy()

	// 启动 Gin 服务
	r.Run(fmt.Sprintf(":%s", config.Get().Http.Port))
}
