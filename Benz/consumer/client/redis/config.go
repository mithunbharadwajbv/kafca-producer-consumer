package redis

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

var (
	client *redis.Client

	ctx = context.Background()
)

func init() {
	//Initializing redis
	client = redis.NewClient(&redis.Options{})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
