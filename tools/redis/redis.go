package redis

import (
	"context"
	"github.com/Heqiaomu/goweb/config"
	"github.com/Heqiaomu/goweb/tools/logger"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"sync"
	"time"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

var redisClient *redis.Client
var once sync.Once

func NewRedisClient(redisCfg config.Redis) (*redis.Client, error) {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password, // no password set
			DB:       redisCfg.DB,       // use default DB
		})

	})
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Errorf("redis connect ping failed, err: %+v", err)
		return nil, err
	}
	logger.Infof("redis connect ping response: pong: [%s]", pong)

	return redisClient, nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := redisClient.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		logger.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := redisClient.Get(rs.Context, key).Result()
	if err != nil {
		logger.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := redisClient.Del(rs.Context, key).Err()
		if err != nil {
			logger.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
