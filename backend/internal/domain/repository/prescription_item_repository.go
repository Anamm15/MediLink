package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PrescriptionItemRepository interface {
	Add(ctx context.Context, data *entity.PrescriptionItem) error
	Update(ctx context.Context, id uuid.UUID, quantity int) error
	Delete(ctx context.Context, id uuid.UUID) error
}
