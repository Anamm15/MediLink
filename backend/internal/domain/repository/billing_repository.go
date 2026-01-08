package repository

import (
	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillingRepository interface {
	GetByID(tx *gorm.DB, billingID uuid.UUID) (entity.Billing, error)
	Create(tx *gorm.DB, billing *entity.Billing) error
}
