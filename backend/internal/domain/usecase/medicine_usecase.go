package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type MedicineUsecase interface {
	GetAll(ctx context.Context, page int) ([]dto.MedicineResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.MedicineResponse, error)
	Search(ctx context.Context, name string, page int) ([]dto.MedicineResponse, error)
	Create(ctx context.Context, request dto.MedicineCreate) (dto.MedicineResponse, error)
	Update(ctx context.Context, id uuid.UUID, request *dto.MedicineUpdate) (dto.MedicineResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
