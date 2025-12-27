package usecase

import (
	"context"
	"time"
)

type NotificationUsecase interface {
	SendOTP(ctx context.Context, key string, email string, identity string, expiration time.Duration) error
}
