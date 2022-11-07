package startup

import (
	"github.com/sun-iot/goweb/config"
	"github.com/sun-iot/goweb/tools/redis"
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
