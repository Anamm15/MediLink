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

func (pmr *PrescriptionItemRepository) Add(ctx context.Context, data *entity.PrescriptionItem) error {
	if err := pmr.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (pmr *PrescriptionItemRepository) Update(ctx context.Context, id uuid.UUID, quantity int) error {
	if err := pmr.db.WithContext(ctx).Model(&entity.PrescriptionItem{}).
		Where("id = ?", id).
		Update("quantity", quantity).Error; err != nil {
		return err
	}
	return nil
}

func (pmr *PrescriptionItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := pmr.db.WithContext(ctx).
		Delete(&entity.PrescriptionItem{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
