// Package server
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/12
 */
package server

import (
	"fmt"
	"github.com/Heqiaomu/goweb/config"
	"github.com/Heqiaomu/goweb/server/router"
	"github.com/Heqiaomu/goweb/tools/logger"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func Start() {
	routers := router.Routers()
	address := fmt.Sprintf(":%d", config.GetConfig().System.Addr)
	s := initServer(address, routers)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	tls := config.GetConfig().System.TLS
	if tls.Enabled {
		logger.Error(s.ListenAndServeTLS(tls.Cert, tls.Key).Error())
	} else {
		logger.Error(s.ListenAndServe().Error())
	}
}

type server interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
