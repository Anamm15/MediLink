package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type PrescriptionUsecase interface {
	GetByPatient(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponseDTO, error)
	GetByDoctor(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponseDTO, error)
	GetDetailByID(ctx context.Context, id uuid.UUID) (*dto.PrescriptionResponseDTO, error)
	Create(ctx context.Context, dto *dto.PrescriptionCreateDTO) (dto.PrescriptionResponseDTO, error)
	Update(ctx context.Context, id uuid.UUID, dto *dto.PrescriptionUpdateDTO) (dto.PrescriptionResponseDTO, error)
	Delete(ctx context.Context, id uuid.UUID) error
	AddMedicine(ctx context.Context, prescriptionID uuid.UUID, data *dto.PrescriptionItemCreateDTO) (dto.PrescriptionItemResponseDTO, error)
	UpdateMedicine(ctx context.Context, id uuid.UUID, quantity int) error
	RemoveMedicine(ctx context.Context, id uuid.UUID) error
}
