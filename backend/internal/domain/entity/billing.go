package entity

import (
	"time"

	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type Billing struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	PatientID     uuid.UUID  `gorm:"type:uuid;not null"`
	AppointmentID *uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	Amount float64 `gorm:"type:numeric(12,2);not null"`

	Status constants.PaymentStatus `gorm:"type:payment_status;default:'unpaid';not null"`

	DueDate     *time.Time `gorm:"type:date"`
	Description *string    `gorm:"type:text"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
