// Package router
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/12
 */
package router

import (
	"github.com/Heqiaomu/goweb/config"
	"github.com/Heqiaomu/goweb/server/middleware"
	"github.com/Heqiaomu/goweb/tools/logger"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	if config.GetConfig().System.TLS.Enabled {
		engine.Use(middleware.LoadTls())
	}
	engine.Use(middleware.Cors()) // 直接放行全部跨域请求

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	public := engine.Group("")
	logger.Infof("Start register common routers")

	commonRouter.InitCommonRouter(public)

	private := engine.Group("")
	private.Use(middleware.JWTAuth())
	{
		userRouter.InitUserRouter(private) // 注册功能api路由
	}
	return engine
}
