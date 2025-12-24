package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type MedicineRepository interface {
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Medicine, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Medicine, error)
	Search(ctx context.Context, name string, limit int, offset int) ([]entity.Medicine, error)
	Create(ctx context.Context, medicine *entity.Medicine) error
	Update(ctx context.Context, medicine *entity.Medicine) error
	Delete(ctx context.Context, id uuid.UUID) error
}
