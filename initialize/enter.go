package initialize

import (
	"Project001/config"
	"Project001/global"

	"github.com/gin-gonic/gin"
)

func GlobalInit() *gin.Engine {
	// 配置文件初始化
	global.Config = config.InitLoadConfig()
	// Log初始化
	// global.Log = logger.NewZapLogCenter(logger.NewZapConfig(
	// 	logger.WithServiceName("takeout"),
	// ))
	// Gorm初始化
	global.DB = InitDatabase(global.Config.DataSource.Dsn(), global.Config.DataSource.DBType)
	// Redis初始化
	global.Redis = initRedis()
	//Router初始化
	router := routerInit()
	return router
}
