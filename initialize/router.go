package initialize

import (
	"Project001/internal/router"

	"github.com/gin-gonic/gin"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter

	// 链路追踪日志中间件
	// r.Use(logger.GinMiddleware(global.Log, "takeout"))

	// admin
	admin := r.Group("/admin")
	{
		allRouter.CommonRouter.InitApiRouter(admin)
	}
	return r
}
