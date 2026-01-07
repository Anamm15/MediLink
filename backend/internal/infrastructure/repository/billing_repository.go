package repository

import (
	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

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

func (r *BillingRepository) Create(tx *gorm.DB, billing *entity.Billing) error {
	if err := tx.Create(billing).Error; err != nil {
		return err
	}
	return nil
}
