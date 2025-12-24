package entity

import (
	"time"

	"github.com/google/uuid"
)

type Medicine struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `gorm:"type:varchar(255);not null"`
	GenericName  *string   `gorm:"type:varchar(255)"`
	Description  *string   `gorm:"type:text"`
	Category     *string   `gorm:"type:varchar(100)"`
	Manufacturer *string   `gorm:"type:varchar(100)"`

	BasePrice              float64 `gorm:"type:numeric(12,2);not null"`
	IsPrescriptionRequired bool    `gorm:"default:false"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
