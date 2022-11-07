package handler

import (
	"github.com/gin-gonic/gin"
	ul "github.com/sun-iot/goweb/logic"
)

type UserApi struct {
}

var logic = new(ul.UserLogic)

func (ua *UserApi) CreateUser(ctx *gin.Context) {
	logic.CreateUser()
}
