package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type DoctorUsecase interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (dto.DoctorProfileResponse, error)
	Find(ctx context.Context, name string, page int) ([]dto.DoctorProfileResponse, error)
	Update(ctx context.Context, userID uuid.UUID, request dto.DoctorUpdateRequest) error
	AddSchedule(ctx context.Context, userID uuid.UUID, request dto.DoctorCreateScheduleRequest) (dto.DoctorScheduleResponse, error)
	UpdateSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, request dto.DoctorUpdateScheduleRequest) error
}
