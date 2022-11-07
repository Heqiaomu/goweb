package handler

import (
	ul "github.com/Heqiaomu/goweb/logic"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

var logic = new(ul.UserLogic)

func (ua *UserApi) CreateUser(ctx *gin.Context) {
	logic.CreateUser()
}
