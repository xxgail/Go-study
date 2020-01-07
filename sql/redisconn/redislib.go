package redisconn

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
)

func InitNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	pong, err := client.Ping().Result()
	fmt.Println("初始化Redis....", pong, err)
}

func GetClient() (c *redis.Client) {
	return client
}
