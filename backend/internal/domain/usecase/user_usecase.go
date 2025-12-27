package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type UserUsecase interface {
	GetAll(ctx context.Context, page int) ([]dto.UserResponseDTO, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (dto.UserProfileResponseDTO, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, data dto.UserUpdateProfileRequestDTO) error
	Delete(ctx context.Context, userID uuid.UUID) error
	SendVerificationUser(ctx context.Context, userID uuid.UUID) error
	VerifyUser(ctx context.Context, userID uuid.UUID, otp string) error
	OnBoardPatient(ctx context.Context, userID uuid.UUID, data dto.PatientCreateRequestDTO) error
	// ApplyAsDoctor(ctx context.Context, userID uuid.UUID, licenseNumber, specialization string) error
	// ApplyAsStaff(ctx context.Context, userID uuid.UUID) error
}
