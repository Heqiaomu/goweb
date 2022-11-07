package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/sun-iot/goweb/config"
	"github.com/sun-iot/goweb/model/common/response"
	"github.com/sun-iot/goweb/model/dbmodel"
	"github.com/sun-iot/goweb/service"
	jwtTool "github.com/sun-iot/goweb/tools/jwt"
	"github.com/sun-iot/goweb/tools/logger"
	"github.com/sun-iot/goweb/tools/timer"
	"go.uber.org/zap"
	"time"
)

var jwtService = new(service.JwtService)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中
		// 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		// 检查 Token 是否在黑名单，或者是已经失效
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := jwtTool.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwtTool.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Sub(time.Now()) < time.Duration(claims.BufferTime) {
			dr, _ := timer.ParseDuration(config.GetConfig().JWT.ExpiresTime)
			claims.ExpiresAt = jwtTool.ParseNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", cast.ToString(newClaims.ExpiresAt.Unix()))
			if config.GetConfig().System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					logger.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(dbmodel.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
