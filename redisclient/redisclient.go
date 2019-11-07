package redisclient

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gitlab.com/ironcore864/actsctrl/config"
)

var client *redis.Client

func getClient() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:        fmt.Sprintf("%s:%d", config.Conf.RedisHost, config.Conf.RedisPort),
			Password:    "", // no password set
			DB:          0,  // use default DB
			DialTimeout: 3 * time.Second,
			PoolSize:    300,
		})
	}
	return client
}

// Exists test if a key exists
func Exists(key string) (int64, error) {
	return getClient().Exists(key).Result()
}

// Incr add a key's value by 1
func Incr(key string) (int64, error) {
	return getClient().Incr(key).Result()
}

// Set a key
func Set(key string, data int, expiration time.Duration) (string, error) {
	return getClient().Set(key, data, expiration).Result()
}
