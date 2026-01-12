package repository

import (
	"context"
	"time"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type DoctorScheduleRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.DoctorSchedule, error)
	GetByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]entity.DoctorSchedule, error)
	GetSchedulesByDate(ctx context.Context, doctorID uuid.UUID, date time.Time) ([]entity.DoctorSchedule, error)
	GetSchedulesByDay(ctx context.Context, doctorID uuid.UUID, day string) ([]entity.DoctorSchedule, error)
	Create(ctx context.Context, schedule *entity.DoctorSchedule) (*entity.DoctorSchedule, error)
	Update(ctx context.Context, schedule *entity.DoctorSchedule) error
	UpdateStatus(ctx context.Context, id uuid.UUID, isActive bool) error
	Delete(ctx context.Context, id uuid.UUID, doctorID uuid.UUID) error
}
