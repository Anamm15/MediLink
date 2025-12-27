package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) repository.RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) GetByTokenHash(ctx context.Context, token string) (*entity.RefreshToken, error) {
	refreshToken := &entity.RefreshToken{}
	if err := r.db.WithContext(ctx).
		Where("token = ?", token).
		First(refreshToken).Error; err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func (r *RefreshTokenRepository) Create(ctx context.Context, RefreshToken *entity.RefreshToken) error {
	if err := r.db.WithContext(ctx).Create(RefreshToken).Error; err != nil {
		return err
	}
	return nil
}

func (r *RefreshTokenRepository) Save(ctx context.Context, token *entity.RefreshToken) error {
	if err := r.db.WithContext(ctx).Save(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *RefreshTokenRepository) RevokeToken(ctx context.Context, token string) error {
	return r.db.
		WithContext(ctx).
		Model(&entity.RefreshToken{}).
		Where("token = ?", token).
		Update("is_revoked", true).
		Error
}

func (r *RefreshTokenRepository) RotateToken(ctx context.Context, oldTokenID uuid.UUID, newToken *entity.RefreshToken) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.RefreshToken{}).Where("id = ?", oldTokenID).Update("is_used", true).Error; err != nil {
			return err
		}
		if err := tx.Create(newToken).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *RefreshTokenRepository) RevokeFamily(ctx context.Context, familyID uuid.UUID) error {
	return r.db.Model(&entity.RefreshToken{}).
		Where("family_id = ?", familyID).
		Update("is_revoked", true).Error
}

func (r *RefreshTokenRepository) Delete(ctx context.Context, userID uuid.UUID, token string) error {
	if err := r.db.WithContext(ctx).
		Delete(&entity.RefreshToken{}, "user_id = ? AND token = ?", userID, token).Error; err != nil {
		return err
	}
	return nil
}
