package redismanager

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func GetRedisClient(redisIpAddress string, password string, dbIndex int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisIpAddress,
		Password: password,
		DB:       dbIndex,
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

func (client *RedisClient) DeleteData(ctx context.Context, key string) (int64, error) {
	val, err := client.Client.Del(ctx, key).Result()
	return val, err
}

func (client *RedisClient) GetAllKeys(ctx context.Context) ([]string, error) {
	// https://redis.uptrace.dev/get-all-keys/
	iter := client.Client.Scan(ctx, 0, "*", 0).Iterator()
	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	err := iter.Err()
	if err != nil {
		panic(err)
	}
	return keys, err
}
