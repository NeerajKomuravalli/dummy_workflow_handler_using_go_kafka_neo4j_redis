package redismanager

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func GetRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisIpAddress,
		Password: Password,
		DB:       DbIndex,
	})
	return &RedisClient{rdb}
}

func (client *RedisClient) PutData(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := client.Client.Set(ctx, key, value, expiration).Err()
	return err
}

func (client *RedisClient) GetData(ctx context.Context, key string) (string, error) {
	val, err := client.Client.Get(ctx, key).Result()
	return val, err
}

func (client *RedisClient) GetAllKeys(ctx context.Context) {
	// https://redis.uptrace.dev/get-all-keys/
	iter := client.Client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println(iter.Val()) // Will print keys one after another
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}
