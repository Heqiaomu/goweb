// Package middleware
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/13
 */
package middleware

import "github.com/gin-gonic/gin"

type Handler struct {
}

var middle []gin.HandlerFunc

func setHandlerFunc(handler ...gin.HandlerFunc) {
	middle = append(middle, handler...)
}

func GetHandlerFunc() []gin.HandlerFunc {
	setHandlerFunc(
		Cors(), 
		GinLogger(false),
		RecoveryWithLogger(false))
	return middle
}
