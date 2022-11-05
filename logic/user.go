package logic

import (
	us "gitee.com/goweb/service"
)

type UserLogic struct {
}

var userService = new(us.UserService)

func (ul *UserLogic) CreateUser() (interface{}, error) {
	return userService.CreateUser()
}
