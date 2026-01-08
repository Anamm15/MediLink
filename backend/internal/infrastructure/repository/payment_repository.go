package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) repository.PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (*entity.Payment, error) {
	var payment entity.Payment
	if err := r.db.WithContext(ctx).First(&payment, "id = ?", paymentID).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) Create(ctx context.Context, payment *entity.Payment) error {
	if err := r.db.WithContext(ctx).Create(payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) UpdateStatus(tx *gorm.DB, billingID uuid.UUID, status enum.PaymentStatus) error {
	if err := tx.Model(&entity.Payment{}).
		Where("billing_id = ?", billingID).
		Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
