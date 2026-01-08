package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *entity.Payment) error
	GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (*entity.Payment, error)
	UpdateStatus(tx *gorm.DB, appointmentID uuid.UUID, status enum.PaymentStatus) error
}
