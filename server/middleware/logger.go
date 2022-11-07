// Package middleware
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/13
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sun-iot/goweb/tools/logger"
	"go.uber.org/zap"
	"time"
)

var GinLogger = func(utc bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		urlPath := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		if utc {
			end = end.UTC()
		}

		if len(ctx.Errors) > 0 {
			for _, e := range ctx.Errors.Errors() {
				logger.Errorf("Manage err when doing gin context. Now to do check. Err: [%v].", e)
			}
		} else {
			logger.Info(urlPath,
				zap.Int("status", ctx.Writer.Status()),
				zap.String("method", ctx.Request.Method),
				zap.String("path", urlPath),
				zap.String("query", query),
				zap.Duration("latency", latency),
			)
		}
	}
}

//
//var RecoveryWithLogger = func(stack bool) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				var brokenPipe bool
//				if ne, ok := err.(*net.OpError); ok {
//					if se, ok := ne.Err.(*os.SyscallError); ok {
//						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
//							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
//							brokenPipe = true
//						}
//					}
//				}
//
//				httpRequest, _ := httputil.DumpRequest(c.Request, true)
//				if brokenPipe {
//					logger.Error(c.Request.URL.Path,
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//					// If the connection is dead, we can't write a status to it.
//					c.Error(err.(error)) // nolint: errcheck
//					c.Abort()
//					return
//				}
//
//				if stack {
//					logger.Error("Manage err when doing Recovery from panic, because panic. Please check.",
//						zap.Time("time", time.Now()),
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//						zap.String("stack", string(debug.Stack())),
//					)
//				} else {
//					logger.Error("Manage err when doing Recovery from panic, because panic. Please check.",
//						zap.Time("time", time.Now()),
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//				}
//				c.AbortWithStatus(http.StatusInternalServerError)
//			}
//		}()
//		c.Next()
//	}
//}
