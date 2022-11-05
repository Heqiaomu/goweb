package router

import "github.com/gin-gonic/gin"

type CommonRouter struct {
}

var commonRouter = new(CommonRouter)

func (cr *CommonRouter) InitCommonRouter(root *gin.RouterGroup) {
	root.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	commonGroup := root.Group("/api/v1/common")
	commonGroup.POST("/home/list", nil)
}
