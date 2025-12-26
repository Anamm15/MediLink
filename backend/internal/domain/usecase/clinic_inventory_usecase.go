package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type ClinicInventoryUsecase interface {
	GetByClinic(ctx context.Context, clinicID uuid.UUID) ([]dto.ClinicInventoryResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.ClinicInventoryResponse, error)
	Create(ctx context.Context, request dto.ClinicInventoryCreateRequest) (dto.ClinicInventoryResponse, error)
	Update(ctx context.Context, id uuid.UUID, request dto.ClinicInventoryUpdateRequest) (dto.ClinicInventoryResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
