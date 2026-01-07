package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *entity.Payment) error
	GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (*entity.Payment, error)
}
