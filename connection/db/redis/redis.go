package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/i-jonathan/pharmacy-api/config"
)

type repo struct {
	Conn *redis.Client
}

func NewRedisConnection() (*repo, error) {
	config2 := config.GetConfig()
	redisRepo := new(repo)

	rdb := redis.NewClient(&redis.Options{
		Addr:     config2.RedisAddress,
		Password: config2.RedisPass,
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.New("Error starting redis connection: " + err.Error())
	}
	redisRepo.Conn = rdb
	return redisRepo, nil
}
