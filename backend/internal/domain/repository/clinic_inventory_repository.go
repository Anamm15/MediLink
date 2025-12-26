package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type ClinicInventoryRepository interface {
	GetAll(ctx context.Context) ([]entity.ClinicInventory, error)
	GetByClinicID(ctx context.Context, clinicID uuid.UUID) ([]entity.ClinicInventory, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.ClinicInventory, error)
	Create(ctx context.Context, clinicInventory *entity.ClinicInventory) error
	Update(ctx context.Context, clinicInventory *entity.ClinicInventory) error
	Delete(ctx context.Context, id uuid.UUID) error
}
