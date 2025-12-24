package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrescriptionMedicineRepository struct {
	db *gorm.DB
}

func NewPrescriptionMedicineRepository(db *gorm.DB) repository.PrescriptionMedicineRepository {
	return &PrescriptionMedicineRepository{
		db: db,
	}
}

func (pmr *PrescriptionMedicineRepository) Add(ctx context.Context, data *entity.PrescriptionMedicine) error {
	return pmr.db.WithContext(ctx).Create(data).Error
}

func (pmr *PrescriptionMedicineRepository) Update(ctx context.Context, id uuid.UUID, quantity int) error {
	return pmr.db.WithContext(ctx).Model(&entity.PrescriptionMedicine{}).Where("id = ?", id).Update("quantity", quantity).Error
}

func (pmr *PrescriptionMedicineRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return pmr.db.WithContext(ctx).Delete(&entity.PrescriptionMedicine{}, "id = ?", id).Error
}
