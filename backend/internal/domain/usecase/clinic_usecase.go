package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type ClinicUsecase interface {
	GetAll(ctx context.Context, page int) ([]dto.ClinicResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.ClinicResponse, error)
	Find(ctx context.Context, name string, page int) ([]dto.ClinicResponse, error)
	Create(ctx context.Context, request dto.ClinicCreateRequest) (dto.ClinicResponse, error)
	Update(ctx context.Context, id uuid.UUID, request dto.ClinicUpdateRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
	AssignDoctor(ctx context.Context, request dto.AssignDoctorRequest) error
	RemoveDoctor(ctx context.Context, request dto.RemoveDoctorRequest) error
}
