package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrescriptionItemRepository struct {
	db *gorm.DB
}

func NewPrescriptionItemRepository(db *gorm.DB) repository.PrescriptionItemRepository {
	return &PrescriptionItemRepository{
		db: db,
	}
}

func (r *PrescriptionItemRepository) Add(ctx context.Context, data *entity.PrescriptionItem) error {
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (r *PrescriptionItemRepository) Update(ctx context.Context, id uuid.UUID, quantity int) error {
	if err := r.db.WithContext(ctx).Model(&entity.PrescriptionItem{}).
		Where("id = ?", id).
		Update("quantity", quantity).Error; err != nil {
		return err
	}
	return nil
}

func (r *PrescriptionItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Delete(&entity.PrescriptionItem{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
