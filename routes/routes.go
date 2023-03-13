package routes

import (
	logger "Dapp/logs"
	"Dapp/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func SetUp(mode string) *gin.Engine {
	//如果设置mode为release则设置gin为该模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middlewares.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//管理后台路由组
	v := r.Group("/api").Use(middlewares.JudgePermissionWeb)
	{
		v.GET("")

	}

	return r
}
