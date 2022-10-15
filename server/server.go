// Package server
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/12
 */
package server

import (
	"gitee.com/goweb/logic/middleware"
	"github.com/gin-gonic/gin"
)

func Server() {
	eg := gin.New()
	eg.Use(middleware.GinLogger(false), middleware.RecoveryWithLogger(false), middleware.Cors())
}

func start() {

}
