package redis_client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository interface {
	Save(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SaveUnlimited(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) Repository {
	return &redisRepository{client: client}
}

func (rr *redisRepository) Save(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	result := rr.client.Set(ctx, key, data, expiration)
	return result.Err()
}

func (rr *redisRepository) SaveUnlimited(ctx context.Context, key string, value interface{}) error {
	return rr.Save(ctx, key, value, 0)
}

func (rr *redisRepository) Get(ctx context.Context, key string) (string, error) {
	value, err := rr.client.Get(ctx, key).Result()
	fmt.Println("TEST GET REDIS : ", value)
	if err != nil {
		return "", err
	}
	return value, err
}
