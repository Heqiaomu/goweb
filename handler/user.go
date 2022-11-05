package handler

import (
	ul "gitee.com/goweb/logic"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

var logic = new(ul.UserLogic)

func (ua *UserApi) CreateUser(ctx *gin.Context) {
	logic.CreateUser()
}
