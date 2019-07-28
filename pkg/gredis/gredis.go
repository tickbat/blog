package gredis

import (
	"blog/pkg/logging"
	"blog/pkg/setting"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

var redisDb *redis.Client

func init() {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     setting.Redis.Host,
		Password: setting.Redis.Password,
		DB:       0,
	})
	_, err := redisDb.Ping().Result()
	if err != nil {
		logging.Fatal("redis error: ", err)
	}
}

func Set(key string, data interface{}, time time.Duration) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := redisDb.Set(key, value, time).Err(); err != nil {
		return err
	}
	return nil
}

func Get(key string) ([]byte, error) {
	val, err := redisDb.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}
