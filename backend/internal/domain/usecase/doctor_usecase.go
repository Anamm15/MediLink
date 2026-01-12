package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type DoctorUsecase interface {
	Me(ctx context.Context, userID uuid.UUID) (dto.DoctorProfileResponse, error)
	GetProfile(ctx context.Context, doctorID uuid.UUID) (dto.DoctorProfileResponse, error)
	Find(ctx context.Context, name string, page int) ([]dto.DoctorProfileResponse, error)
	Update(ctx context.Context, userID uuid.UUID, request dto.DoctorUpdateRequest) error
	GetDoctorSchedules(ctx context.Context, doctorID uuid.UUID) ([]dto.DoctorScheduleResponse, error)
	GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (dto.DoctorScheduleResponse, error)
	GetAvailableSchedules(ctx context.Context, doctorID uuid.UUID, date string, day string) ([]dto.DoctorScheduleResponse, error)
	AddSchedule(ctx context.Context, userID uuid.UUID, request dto.DoctorCreateScheduleRequest) (dto.DoctorScheduleResponse, error)
	UpdateSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, request dto.DoctorUpdateScheduleRequest) error
	UpdateStatusSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, request dto.DoctorUpdateStatusScheduleRequest) error
	DeleteSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID) error
}
