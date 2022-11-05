// Package cmd
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/11
 */
package main

import (
	"gitee.com/goweb/server"
	"gitee.com/goweb/server/startup"
	"log"
)

func main() {
	// 初始化系统
	log.Printf("Start to init viper")
	startup.InitViper()
	log.Printf("Start to init config")
	startup.InitConfig()
	log.Printf("Start to init log")
	startup.InitLogger()
	log.Printf("Start to init redis, if enabled")
	startup.InitRedis()
	log.Printf("Start to init mysql, if enabled")
	startup.InitMySQL()
	log.Printf("All ready. Start server")

	server.Start() // 启动 Gin 服务

}
