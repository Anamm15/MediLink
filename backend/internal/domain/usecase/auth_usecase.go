package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type AuthUsecase interface {
	Login(ctx context.Context, req dto.LoginRequest) (string, string, error)
	Register(ctx context.Context, req dto.RegistrationRequest) (dto.RegistrationResponse, error)
	RefreshToken(ctx context.Context, token string) (string, string, error)
	Logout(ctx context.Context, token string) error
	ChangePassword(ctx context.Context, userID uuid.UUID, data dto.ChangePasswordRequest) error
	RequestResetPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, data dto.ResetPasswordRequest) error
}
