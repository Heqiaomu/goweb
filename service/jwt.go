package service

import (
	"context"
	"github.com/Heqiaomu/goweb/config"
	"github.com/Heqiaomu/goweb/model/dbmodel"
	"github.com/Heqiaomu/goweb/tools/cache"
	"github.com/Heqiaomu/goweb/tools/logger"
	"github.com/Heqiaomu/goweb/tools/mysql"
	"github.com/Heqiaomu/goweb/tools/redis"
	timeUtil "github.com/Heqiaomu/goweb/tools/timer"
	"go.uber.org/zap"
)

type JwtService struct {
}

// JsonInBlacklist 拉黑
func (jwtService *JwtService) JsonInBlacklist(jwtList dbmodel.JwtBlacklist) (err error) {
	err = mysql.GetDB(context.Background()).Create(&jwtList).Error
	if err != nil {
		return
	}
	cache.GetLocalCache().SetDefault(jwtList.Jwt, struct{}{})
	return
}

// IsBlacklist 判断是否在黑名单
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := cache.GetLocalCache().Get(jwt)
	return ok
}

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = redis.GetRedisClient().Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := timeUtil.ParseDuration(config.GetConfig().JWT.ExpiresTime)
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
		cache.GetLocalCache().SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
