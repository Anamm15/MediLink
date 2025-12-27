package usecase

import (
	"context"
	"fmt"
	"time"

	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/infrastructure/mail"
	"MediLink/internal/utils"
)

type NotificationUsecase struct {
	cacheRepository repository.CacheRepository
}

func NewNotificationUsecase(cacheRepository repository.CacheRepository) usecase.NotificationUsecase {
	return &NotificationUsecase{
		cacheRepository: cacheRepository,
	}
}

func (u *NotificationUsecase) SendOTP(ctx context.Context, key string, email string, identity string, expiration time.Duration) error {
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return err
	}

	err = u.cacheRepository.Set(ctx, key, otp, expiration)
	if err != nil {
		return err
	}

	emailBody := utils.BuildOTPEmailBody(identity, otp)

	go func() {
		err := mail.SendEmail(email, "Kode Verifikasi Keamanan - MediLink", emailBody)
		if err != nil {
			fmt.Printf("Gagal mengirim email ke %s: %v\n", email, err)
		}
	}()
	return nil
}
