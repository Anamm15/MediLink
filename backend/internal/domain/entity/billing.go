package entity

import (
	"time"

	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type Billing struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PatientID     uuid.UUID  `gorm:"type:uuid;not null"`
	AppointmentID *uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	TotalAmount float64            `gorm:"type:numeric(12,2);not null"`
	Status      enum.PaymentStatus `gorm:"type:varchar(20);default:'unpaid';not null"`
	IssuedAt    time.Time          `gorm:"type:timestamptz;default:now();not null"`
	PaidAt      *time.Time         `gorm:"type:timestamptz"`
}
