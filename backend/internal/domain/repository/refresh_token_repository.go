package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type RefreshTokenRepository interface {
	GetByTokenHash(ctx context.Context, token string) (*entity.RefreshToken, error)
	Create(ctx context.Context, refreshToken *entity.RefreshToken) error
	Save(ctx context.Context, refreshToken *entity.RefreshToken) error
	RevokeToken(ctx context.Context, token string) error
	RotateToken(ctx context.Context, oldTokenID uuid.UUID, newToken *entity.RefreshToken) error
	RevokeFamily(ctx context.Context, familyID uuid.UUID) error
	Delete(ctx context.Context, userID uuid.UUID, token string) error
}
