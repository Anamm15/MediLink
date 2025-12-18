package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type DoctorScheduleRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.DoctorSchedule, error)
	GetByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]entity.DoctorSchedule, error)
	Create(ctx context.Context, schedule *entity.DoctorSchedule) (*entity.DoctorSchedule, error)
	Update(ctx context.Context, schedule *entity.DoctorSchedule) error
	Delete(ctx context.Context, id uuid.UUID) error
}
