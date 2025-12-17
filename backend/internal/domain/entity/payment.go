package entity

import (
	"time"

	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type Payment struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	BillingID uuid.UUID `gorm:"type:uuid;not null"`

	PaymentMethod string  `gorm:"type:varchar(50)"`
	Amount        float64 `gorm:"type:numeric(12,2);not null"`

	Status constants.PaymentStatus `gorm:"type:payment_status;not null"`

	PaymentGatewayOrderID       *string `gorm:"type:varchar(255)"`
	PaymentGatewayTransactionID *string `gorm:"type:varchar(255)"`
	PaymentURL                  *string `gorm:"type:text"`
	VANumber                    *string `gorm:"type:varchar(100)"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
