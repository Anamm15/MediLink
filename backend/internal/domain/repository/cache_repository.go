package repository

import (
	"context"
	"time"
)

type CacheRepository interface {
	// Set menyimpan data dengan durasi expired (TTL)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	// Get mengambil data string
	Get(ctx context.Context, key string) (string, error)
	// Delete menghapus data
	Delete(ctx context.Context, key string) error
}
