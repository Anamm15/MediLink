package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type PrescriptionUsecase interface {
	GetByPatient(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponse, error)
	GetByDoctor(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponse, error)
	GetDetailByID(ctx context.Context, id uuid.UUID) (*dto.PrescriptionResponse, error)
	Create(ctx context.Context, dto *dto.PrescriptionCreate) (dto.PrescriptionResponse, error)
	Update(ctx context.Context, id uuid.UUID, dto *dto.PrescriptionUpdate) (dto.PrescriptionResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
	AddMedicine(ctx context.Context, prescriptionID uuid.UUID, data *dto.PrescriptionItemCreate) (dto.PrescriptionItemResponse, error)
	UpdateMedicine(ctx context.Context, id uuid.UUID, quantity int) error
	RemoveMedicine(ctx context.Context, id uuid.UUID) error
}
