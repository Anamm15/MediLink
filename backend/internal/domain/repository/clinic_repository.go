package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type ClinicRepository interface {
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Clinic, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Clinic, error)
	Find(ctx context.Context, name string, limit int, offset int) ([]entity.Clinic, int64, error)
	Create(ctx context.Context, clinic *entity.Clinic) (*entity.Clinic, error)
	Update(ctx context.Context, clinic *entity.Clinic) error
	Delete(ctx context.Context, id uuid.UUID) error
}
