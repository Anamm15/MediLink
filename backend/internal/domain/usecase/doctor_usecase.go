package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type DoctorUsecase interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (dto.DoctorProfileResponse, error)
	Find(ctx context.Context, name string, page int) ([]dto.DoctorProfileResponse, error)
	Update(ctx context.Context, userID uuid.UUID, doctorID uuid.UUID, data dto.DoctorUpdateRequest) error
	AddSchedule(ctx context.Context, data dto.DoctorCreateScheduleRequest) (dto.DoctorScheduleResponse, error)
	UpdateSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, data dto.DoctorUpdateScheduleRequest) error
}
