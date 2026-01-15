package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type MedicineUsecase interface {
	GetAll(ctx context.Context, page string, limit string) (dto.MedicineSearchResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.MedicineResponse, error)
	Search(ctx context.Context, name string, page string, limit string) (dto.MedicineSearchResponse, error)
	Create(ctx context.Context, request dto.MedicineCreate) (dto.MedicineResponse, error)
	Update(ctx context.Context, id uuid.UUID, request *dto.MedicineUpdate) (dto.MedicineResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
