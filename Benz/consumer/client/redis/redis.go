package redis

import (
	"log"
	"time"
)

var (
	expiry = 24 * time.Hour
)

type RedisClient interface {
	SetDataWithExpiry(key string, value string) error
	GetValue(key string) (string, error)
}

type redisClient struct{}

func GetNewRedisClient() RedisClient {
	return &redisClient{}
}

func (redisClient *redisClient) SetDataWithExpiry(key string, value string) error {
	status := client.Set(ctx, key, value, expiry)
	if status.Err() != nil {
		log.Println("Error while Setting the data")
		return status.Err()
	}
	return nil
}

func (redisClient *redisClient) GetValue(key string) (string, error) {
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Error while Getting the data")
		return "", err
	}
	return value, nil

}
