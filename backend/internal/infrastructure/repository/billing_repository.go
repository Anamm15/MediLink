package repository

import (
	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillingRepository struct {
	db *gorm.DB
}

func NewBillingRepository(db *gorm.DB) repository.BillingRepository {
	return &BillingRepository{
		db: db,
	}
}

func (r *BillingRepository) GetByID(tx *gorm.DB, billingID uuid.UUID) (entity.Billing, error) {
	var billing entity.Billing
	if err := tx.First(&billing, "id = ?", billingID).Error; err != nil {
		return entity.Billing{}, err
	}
	return billing, nil
}

func (r *BillingRepository) Create(tx *gorm.DB, billing *entity.Billing) error {
	if err := tx.Create(billing).Error; err != nil {
		return err
	}
	return nil
}
