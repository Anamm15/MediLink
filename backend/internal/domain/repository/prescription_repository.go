package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PrescriptionRepository interface {
	GetByPatient(ctx context.Context, patientID uuid.UUID) ([]entity.Prescription, error)
	GetByDoctor(ctx context.Context, doctorID uuid.UUID) ([]entity.Prescription, error)
	GetDetailByID(ctx context.Context, id uuid.UUID) (*entity.Prescription, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Prescription, error)
	Create(ctx context.Context, prescription *entity.Prescription) error
	Update(ctx context.Context, prescription *entity.Prescription) error
	Delete(ctx context.Context, id uuid.UUID) error
}
