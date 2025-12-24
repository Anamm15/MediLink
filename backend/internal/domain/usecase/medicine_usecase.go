package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type MedicineUsecase interface {
	GetAll(ctx context.Context, page int) ([]dto.MedicineResponseDTO, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.MedicineResponseDTO, error)
	Search(ctx context.Context, name string, page int) ([]dto.MedicineResponseDTO, error)
	Create(ctx context.Context, medicine dto.MedicineCreateDTO) (dto.MedicineResponseDTO, error)
	Update(ctx context.Context, id uuid.UUID, medicine *dto.MedicineUpdateDTO) (dto.MedicineResponseDTO, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
