package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type ClinicUsecase interface {
	GetAll(ctx context.Context, page int) ([]dto.ClinicResponseDTO, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.ClinicResponseDTO, error)
	Find(ctx context.Context, name string, page int) ([]dto.ClinicResponseDTO, error)
	Create(ctx context.Context, data dto.ClinicCreateRequestDTO) (dto.ClinicResponseDTO, error)
	Update(ctx context.Context, id uuid.UUID, data dto.ClinicUpdateRequestDTO) error
	Delete(ctx context.Context, id uuid.UUID) error
	AssignDoctor(ctx context.Context, data dto.AssignDoctorRequest) error
	RemoveDoctor(ctx context.Context, data dto.RemoveDoctorRequest) error
}
