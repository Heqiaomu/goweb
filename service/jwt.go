package service

import (
	"context"
	"gitee.com/goweb/config"
	"gitee.com/goweb/model/dbmodel"
	"gitee.com/goweb/tools/local_cache"
	"gitee.com/goweb/tools/logger"
	"gitee.com/goweb/tools/mysql"
	"gitee.com/goweb/tools/redis"
	timer2 "gitee.com/goweb/tools/timer"
	"go.uber.org/zap"
)

type JwtService struct {
}

//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList dbmodel.JwtBlacklist) (err error) {
	err = mysql.GetDB(context.Background()).Create(&jwtList).Error
	if err != nil {
		return
	}
	local_cache.GetLocalCache().SetDefault(jwtList.Jwt, struct{}{})
	return
}

// @function: IsBlacklist
// @description: 判断JWT是否在黑名单内部

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := local_cache.GetLocalCache().Get(jwt)
	return ok
}

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = redis.GetRedisClient().Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := timer2.ParseDuration(config.GetConfig().JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = redis.GetRedisClient().Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := mysql.GetDB(context.Background()).Model(&dbmodel.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		logger.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		local_cache.GetLocalCache().SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
