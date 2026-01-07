package repository

import (
	"MediLink/internal/domain/entity"

	"gorm.io/gorm"
)

type BillingRepository interface {
	Create(tx *gorm.DB, billing *entity.Billing) error
}
