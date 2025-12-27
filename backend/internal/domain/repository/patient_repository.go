package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PatientRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Patient, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Patient, error)
	Create(ctx context.Context, patient *entity.Patient) (*entity.Patient, error)
	Update(ctx context.Context, patient *entity.Patient) error
}
