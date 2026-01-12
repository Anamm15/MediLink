package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type DoctorRepository interface {
	GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.Doctor, error)
	GetProfileByID(context.Context, uuid.UUID) (*entity.Doctor, error)
	GetByID(ctx context.Context, doctorID uuid.UUID) (*entity.Doctor, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Doctor, error)
	Find(ctx context.Context, name string, limit int, offset int) ([]entity.Doctor, error)
	Create(ctx context.Context, doctor *entity.Doctor) (*entity.Doctor, error)
	Update(ctx context.Context, doctor *entity.Doctor) error
	Delete(ctx context.Context, id uuid.UUID) error
}
