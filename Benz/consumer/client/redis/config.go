package redis

import (
	"consumer/config"
	"context"

	redis "github.com/go-redis/redis/v8"
)

var (
	client *redis.Client

	ctx       = context.Background()
	radisPort = config.Conf.RadisPort
)

func init() {
	//Initializing redis
	client = redis.NewClient(&redis.Options{Addr: radisPort})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
