package repository

import (
	"context"
	"errors"
	"time"

	domainRepo "MediLink/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type redisRepository struct {
	rdb *redis.Client
}

// Constructor
func NewRedisRepository(rdb *redis.Client) domainRepo.CacheRepository {
	return &redisRepository{rdb: rdb}
}

func (r *redisRepository) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	// SetNX (Set if Not Exists) atau Set biasa. Untuk OTP, Set biasa oke.
	return r.rdb.Set(ctx, key, value, ttl).Err()
}

func (r *redisRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// Ubah error "redis: nil" menjadi error umum "data not found"
			// agar usecase tidak perlu import package redis
			return "", errors.New("data not found")
		}
		return "", err
	}
	return val, nil
}

func (r *redisRepository) Delete(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}
