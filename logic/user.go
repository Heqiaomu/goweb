package logic

import (
	us "github.com/Heqiaomu/goweb/service"
)

type UserLogic struct {
}

var userService = new(us.UserService)

func (ul *UserLogic) CreateUser() (interface{}, error) {
	return userService.CreateUser()
}
