// Package cmd
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/11
 */
package main

import (
	"github.com/Heqiaomu/goweb/server"
	"github.com/Heqiaomu/goweb/server/startup"
)

func main() {
	// 初始化系统
	startup.InitServer()
	server.Start() // 启动 Gin 服务
}
