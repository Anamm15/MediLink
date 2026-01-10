package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type UserUsecase interface {
	GetAll(ctx context.Context, page int) ([]dto.UserResponse, error)
	Me(ctx context.Context, userID uuid.UUID) (dto.UserResponse, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (dto.UserProfileResponse, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, request dto.UserUpdateProfileRequest) error
	Delete(ctx context.Context, userID uuid.UUID) error
	SendVerificationUser(ctx context.Context, userID uuid.UUID) error
	VerifyUser(ctx context.Context, userID uuid.UUID, otp string) error
	OnBoardPatient(ctx context.Context, userID uuid.UUID, request dto.PatientCreateRequest) (dto.PatientResponse, error)
	// ApplyAsDoctor(ctx context.Context, userID uuid.UUID, licenseNumber, specialization string) error
	// ApplyAsStaff(ctx context.Context, userID uuid.UUID) error
}
