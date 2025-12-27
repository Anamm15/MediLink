package usecase

import (
	"context"
	"errors"
	"time"

	"MediLink/internal/domain/entity"
	errs "MediLink/internal/domain/errors"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/utils"

	"github.com/google/uuid"
)

type AuthUsecase struct {
	authRepository      repository.RefreshTokenRepository
	userRepository      repository.UserRepository
	cacheRepository     repository.CacheRepository
	notificationUsecase usecase.NotificationUsecase
}

func NewAuthUsecase(
	authRepository repository.RefreshTokenRepository,
	userRepository repository.UserRepository,
	cacheRepository repository.CacheRepository,
	notificationUsecase usecase.NotificationUsecase,
) usecase.AuthUsecase {
	return &AuthUsecase{
		authRepository:      authRepository,
		userRepository:      userRepository,
		cacheRepository:     cacheRepository,
		notificationUsecase: notificationUsecase,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, request dto.LoginRequest) (string, string, error) {
	user, err := u.userRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		return "", "", errs.ErrEmailOrPass
	}

	if err := utils.ComparePassword(user.Password, request.Password); err != nil {
		return "", "", errs.ErrEmailOrPass
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshTokenStr, err := utils.GenerateRandomString(32)
	if err != nil {
		return "", "", err
	}

	tokenID := uuid.New()

	refreshTokenRecord := entity.RefreshToken{
		ID:        tokenID,
		UserID:    user.ID,
		Token:     refreshTokenStr,
		FamilyID:  tokenID,
		IsUsed:    false,
		IsRevoked: false,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := u.authRepository.Create(ctx, &refreshTokenRecord); err != nil {
		return "", "", err
	}

	return accessToken, refreshTokenStr, nil
}

func (u *AuthUsecase) Register(ctx context.Context, request dto.RegistrationRequest) (dto.RegistrationResponse, error) {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return dto.RegistrationResponse{}, err
	}

	user := &entity.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    hashedPassword,
	}
	createdUser, err := u.userRepository.Create(ctx, user)
	if err != nil {
		return dto.RegistrationResponse{}, err
	}

	return dto.RegistrationResponse{
		ID:                  createdUser.ID,
		RegistrationRequest: request,
	}, nil
}

func (u *AuthUsecase) RefreshToken(ctx context.Context, tokenStr string) (string, string, error) {
	storedToken, err := u.authRepository.GetByTokenHash(ctx, tokenStr)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	// --- Reuse Detection ---
	if storedToken.IsUsed {
		_ = u.authRepository.RevokeFamily(ctx, storedToken.FamilyID)
		return "", "", errors.New("refresh token reuse detected: security alert")
	}

	if storedToken.IsRevoked {
		return "", "", errors.New("token revoked")
	}
	if storedToken.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("token expired")
	}

	user, err := u.userRepository.GetByID(ctx, storedToken.UserID)
	if err != nil {
		return "", "", err
	}

	// --- Token Rotation ---
	newAccessToken, err := utils.GenerateJWT(storedToken.UserID, user.Role)
	if err != nil {
		return "", "", err
	}

	newRefreshTokenStr, err := utils.GenerateRandomString(32)
	if err != nil {
		return "", "", err
	}

	newRefreshTokenRecord := entity.RefreshToken{
		ID:        uuid.New(),
		UserID:    storedToken.UserID,
		Token:     newRefreshTokenStr,
		FamilyID:  storedToken.FamilyID,
		IsUsed:    false,
		IsRevoked: false,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = u.authRepository.RotateToken(ctx, storedToken.ID, &newRefreshTokenRecord)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshTokenStr, nil
}

func (u *AuthUsecase) Logout(ctx context.Context, token string) error {
	return u.authRepository.RevokeToken(ctx, token)
}

func (u *AuthUsecase) ChangePassword(ctx context.Context, userID uuid.UUID, request dto.ChangePasswordRequest) error {
	user, err := u.userRepository.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := utils.ComparePassword(user.Password, *request.OldPassword); err != nil {
		return errs.ErrOldPassIncorrect
	}

	newHashedPassword, err := utils.HashPassword(*request.NewPassword)
	if err != nil {
		return err
	}
	user.Password = newHashedPassword

	return u.userRepository.Update(ctx, user)
}

func (u *AuthUsecase) RequestResetPassword(ctx context.Context, email string) error {
	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	key := "reset-password:" + user.Email
	expiration := 8 * time.Minute
	u.notificationUsecase.SendOTP(ctx, key, user.Email, user.Name, expiration)
	return nil
}

func (u *AuthUsecase) ResetPassword(ctx context.Context, request dto.ResetPasswordRequest) error {
	key := "reset-password:" + request.Email
	storedOTP, err := u.cacheRepository.Get(ctx, key)
	if err != nil {
		return errors.New("OTP has expired or does not exist")
	}

	if storedOTP != request.OTP {
		return errors.New("Invalid OTP code")
	}

	user, err := u.userRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	newHashedPassword, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		return err
	}
	user.Password = newHashedPassword

	return u.userRepository.Update(ctx, user)
}
