package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PrescriptionMedicineRepository interface {
	Add(ctx context.Context, data *entity.PrescriptionMedicine) error
	Update(ctx context.Context, id uuid.UUID, quantity int) error
	Delete(ctx context.Context, id uuid.UUID) error
}
