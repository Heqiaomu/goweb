package startup

import (
	"gitee.com/goweb/config"
	"gitee.com/goweb/tools/redis"
	"log"
)

func InitRedis() error {
	redisConfig := config.GetConfig().Redis
	if redisConfig.Enabled {
		client, err := redis.NewRedisClient(redisConfig)
		if client == nil && err != nil {
			log.Fatalf("Redsi init failed. %s", err.Error())
			return err
		}
	}
	return nil
}
