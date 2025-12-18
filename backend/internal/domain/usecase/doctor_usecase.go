package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type DoctorUsecase interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (dto.DoctorProfileResponseDTO, error)
	Find(ctx context.Context, name string, page int) ([]dto.DoctorProfileResponseDTO, error)
	Update(ctx context.Context, userID uuid.UUID, doctorID uuid.UUID, data dto.DoctorUpdateRequestDTO) error
	AddSchedule(ctx context.Context, userID uuid.UUID, data dto.DoctorCreateScheduleRequestDTO) (dto.DoctorScheduleResponseDTO, error)
	UpdateSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, data dto.DoctorUpdateScheduleRequestDTO) error
}
