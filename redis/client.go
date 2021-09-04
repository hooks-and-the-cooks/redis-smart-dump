package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func InitRedisClient(host string, port string) *redis.Client {
	config := &redis.Options{
		Addr:     host + ":" + port,
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.dbIndex"),
	}
	client := redis.NewClient(config)
	return client
}
