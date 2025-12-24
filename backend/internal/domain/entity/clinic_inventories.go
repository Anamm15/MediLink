package entity

import (
	"github.com/google/uuid"
)

type ClinicInventory struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;index:idx_inventory_clinic_catalog"`
	ClinicID   uuid.UUID `gorm:"type:uuid;not null;index:idx_inventory_clinic_catalog"`
	MedicineID uuid.UUID `gorm:"type:uuid;not null"`

	CurrentStock      int     `gorm:"not null;default:0;check:current_stock >= 0"`
	LowStockThreshold int     `gorm:"default:10"`
	Price             float64 `gorm:"type:numeric(12,2);not null"`

	BatchNumber *string `gorm:"type:varchar(100);"`
	ExpiryDate  *string `gorm:"type:date;"`

	updatedAt int64 `gorm:"autoUpdateTime"`
}
