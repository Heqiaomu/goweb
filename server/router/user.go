package router

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

var userRouter = new(UserRouter)

func (ur *UserRouter) InitUserRouter(root *gin.RouterGroup) {
	userGroup := root.Group("/api/v1/user")

	userGroup.POST("create-user", nil)
}
