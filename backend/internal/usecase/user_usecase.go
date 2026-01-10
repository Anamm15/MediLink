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
	"MediLink/internal/helpers/constants"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo            repository.UserRepository
	patientRepo         repository.PatientRepository
	cacheRepo           repository.CacheRepository
	notificationUsecase usecase.NotificationUsecase
}

func NewUserUsecase(
	userRepo repository.UserRepository,
	patientRepo repository.PatientRepository,
	cacheRepo repository.CacheRepository,
	notificationUsecase usecase.NotificationUsecase,
) usecase.UserUsecase {
	return &userUsecase{
		userRepo:            userRepo,
		patientRepo:         patientRepo,
		cacheRepo:           cacheRepo,
		notificationUsecase: notificationUsecase,
	}
}

func (u *userUsecase) GetAll(ctx context.Context, page int) ([]dto.UserResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	users, err := u.userRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var userDTOs []dto.UserResponse
	for _, user := range users {
		userDTOs = append(userDTOs, dto.ToUserResponse(&user))
	}
	return userDTOs, nil
}

func (u *userUsecase) Me(ctx context.Context, userID uuid.UUID) (dto.UserResponse, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.ToUserResponse(user), nil
}

func (u *userUsecase) GetProfile(ctx context.Context, userID uuid.UUID) (dto.UserProfileResponse, error) {
	var user *entity.User
	var patient *entity.Patient

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		user, err = u.userRepo.GetByID(ctx, userID)
		return err
	})

	// Patient Opsional
	g.Go(func() error {
		var err error
		patient, err = u.patientRepo.GetByUserID(ctx, userID)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	})

	if err := g.Wait(); err != nil {
		return dto.UserProfileResponse{}, err
	}

	response := dto.UserProfileResponse{
		User: dto.ToUserResponse(user),
	}

	// Hanya mapping patient jika datanya memang ditemukan
	if patient != nil {
		patientDTO := dto.ToPatientResponse(patient)
		response.Patient = &patientDTO
	}

	return response, nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, userID uuid.UUID, request dto.UserUpdateProfileRequest) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	request.ToModel(user)
	return u.userRepo.Update(ctx, user)
}

func (u *userUsecase) Delete(ctx context.Context, userID uuid.UUID) error {
	return u.userRepo.Delete(ctx, userID)
}

func (u *userUsecase) SendVerificationUser(ctx context.Context, userID uuid.UUID) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	key := "verify-user-otp:" + user.Email
	expiration := 3 * time.Minute
	u.notificationUsecase.SendOTP(ctx, key, user.Email, user.Name, expiration)
	return nil
}

func (u *userUsecase) VerifyUser(ctx context.Context, userID uuid.UUID, inputOTP string) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return nil
	}

	key := "verify-user-otp:" + user.Email
	storedOTP, err := u.cacheRepo.Get(ctx, key)
	if err != nil {
		return errors.New("OTP has expired or does not exist")
	}

	if storedOTP != inputOTP {
		return errors.New("Invalid OTP code")
	}

	_ = u.cacheRepo.Delete(ctx, key)
	user.IsVerified = true

	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) OnBoardPatient(ctx context.Context, userID uuid.UUID, request dto.PatientCreateRequest) (dto.PatientResponse, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return dto.PatientResponse{}, err
	}

	if !user.IsVerified {
		return dto.PatientResponse{}, errs.ErrUserNotVerified
	}

	var patient entity.Patient
	patient.UserID = userID
	request.ToModel(&patient)
	_, err = u.patientRepo.Create(ctx, &patient)
	if err != nil {
		return dto.PatientResponse{}, err
	}

	user.Role = enum.UserRole(enum.RolePatient)
	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return dto.PatientResponse{}, err
	}

	patientResponse := dto.ToPatientResponse(&patient)
	return patientResponse, nil
}
