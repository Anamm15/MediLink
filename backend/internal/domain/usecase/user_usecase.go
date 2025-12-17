package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type UserUsecase interface {
	Register(ctx context.Context, data dto.UserRegistrationRequestDTO) (dto.UserRegistrationResponseDTO, error)
	Login(ctx context.Context, data dto.UserLoginRequestDTO) (string, error)
	// RefreshToken(ctx context.Context, oldToken string) (string, error)
	// Logout(ctx context.Context, token string) error
	GetAll(ctx context.Context, page int) ([]dto.UserResponseDTO, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (dto.UserResponseDTO, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, data dto.UserUpdateProfileRequestDTO) error
	ChangePassword(ctx context.Context, userID uuid.UUID, data dto.UserChangePasswordRequestDTO) error
	Delete(ctx context.Context, userID uuid.UUID) error
	// OnBoardPatient(ctx context.Context, userID uuid.UUID, medicalHistory string) error
	// ApplyAsDoctor(ctx context.Context, userID uuid.UUID, licenseNumber, specialization string) error
	// ApplyAsStaff(ctx context.Context, userID uuid.UUID) error
}
