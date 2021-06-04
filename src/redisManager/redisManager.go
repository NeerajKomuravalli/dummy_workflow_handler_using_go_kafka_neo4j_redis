package redismanager

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisIpAddress,
		Password: Password, // no password set
		DB:       DbIndex,  // use default DB
	})
	return rdb
}

func PutData(ctx context.Context, client *redis.Client, key string, value interface{}, expiration time.Duration) error {
	err := client.Set(ctx, key, value, expiration).Err()
	return err
}

func GetData(ctx context.Context, client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	return val, err
}
