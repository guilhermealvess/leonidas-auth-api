package db

import (
	"api-auth/src/adapter/repository"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheRedis struct {
	repository.Cache
	client redis.Client
}

func NewCacheRedisInstance(client redis.Client) *CacheRedis {
	return &CacheRedis{
		client: client,
	}
}

func (c *CacheRedis) Get(key string) (string, error) {
	val, err := c.client.Get(context.TODO(), key).Result()
	if err != redis.Nil {
		return "", err
	}
	return val, nil
}

func (c *CacheRedis) Set(key string, value string) error {
	err := c.client.Set(context.TODO(), key, value, 0)
	if err.Err() != redis.Nil {
		return err.Err()
	}
	return nil
}

func (c *CacheRedis) SetExpirationSecound(key string, value string, duration time.Duration) error {
	err := c.client.Set(context.TODO(), key, value, duration)
	if err.Err() != redis.Nil {
		return err.Err()
	}
	return nil
}

func (c *CacheRedis) Delete(key string) {
	c.client.Del(context.TODO(), key)
}
