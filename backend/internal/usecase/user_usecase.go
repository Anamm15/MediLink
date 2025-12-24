package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"MediLink/internal/domain/entity"
	errs "MediLink/internal/domain/errors"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"
	"MediLink/internal/helpers/enum"
	"MediLink/internal/infrastructure/mail"
	"MediLink/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo    repository.UserRepository
	patientRepo repository.PatientRepository
	cacheRepo   repository.CacheRepository
}

func NewUserUsecase(
	userRepo repository.UserRepository,
	patientRepo repository.PatientRepository,
	cacheRepo repository.CacheRepository,
) usecase.UserUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		patientRepo: patientRepo,
		cacheRepo:   cacheRepo,
	}
}

func (u *userUsecase) Register(ctx context.Context, data dto.UserRegistrationRequestDTO) (dto.UserRegistrationResponseDTO, error) {
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return dto.UserRegistrationResponseDTO{}, err
	}

	user := &entity.User{
		Name:        data.Name,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Password:    hashedPassword,
	}
	createdUser, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return dto.UserRegistrationResponseDTO{}, err
	}

	return dto.UserRegistrationResponseDTO{
		ID:                         createdUser.ID,
		UserRegistrationRequestDTO: data,
	}, nil
}

func (u *userUsecase) Login(ctx context.Context, data dto.UserLoginRequestDTO) (string, error) {
	user, err := u.userRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return "", errs.ErrEmailOrPass
	}

	if err := utils.ComparePassword(user.Password, data.Password); err != nil {
		return "", errs.ErrEmailOrPass
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

// func (u *userUsecase) RefreshToken(ctx context.Context, oldToken string) (string, error) {
// 	return u.userRepo.RefreshToken(ctx, oldToken)
// }

// func (u *userUsecase) Logout(ctx context.Context, token string) error {
// 	return u.userRepo.Logout(ctx, token)
// }

func (u *userUsecase) GetAll(ctx context.Context, page int) ([]dto.UserResponseDTO, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	users, err := u.userRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var userDTOs []dto.UserResponseDTO
	for _, user := range users {
		userDTOs = append(userDTOs, dto.MapUserToUserResponseDTO(&user))
	}
	return userDTOs, nil
}

func (u *userUsecase) GetProfile(ctx context.Context, userID uuid.UUID) (dto.UserProfileResponseDTO, error) {
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
		return dto.UserProfileResponseDTO{}, err
	}

	response := dto.UserProfileResponseDTO{
		User: dto.MapUserToUserResponseDTO(user),
	}

	// Hanya mapping patient jika datanya memang ditemukan
	if patient != nil {
		patientDTO := dto.MapPatientToPatientResponseDTO(patient)
		response.Patient = patientDTO
	}

	return response, nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, userID uuid.UUID, data dto.UserUpdateProfileRequestDTO) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	data.AssignToEntity(user)
	return u.userRepo.Update(ctx, user)
}

func (u *userUsecase) ChangePassword(ctx context.Context, userID uuid.UUID, data dto.UserChangePasswordRequestDTO) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := utils.ComparePassword(user.Password, *data.OldPassword); err != nil {
		return errs.ErrOldPassIncorrect
	}

	newHashedPassword, err := utils.HashPassword(*data.NewPassword)
	if err != nil {
		return err
	}
	user.Password = newHashedPassword

	return u.userRepo.Update(ctx, user)
}

func (u *userUsecase) Delete(ctx context.Context, userID uuid.UUID) error {
	return u.userRepo.Delete(ctx, userID)
}

func (u *userUsecase) SendOTP(ctx context.Context, userID uuid.UUID) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return err
	}

	// 3. PENTING: Simpan OTP ke Storage (Redis/DB) dengan Expiration Time
	// Key-nya bisa berupa "otp:userID" atau "otp:email"
	// Contoh menggunakan interface repository cache/redis:
	err = u.cacheRepo.Set(ctx, "otp:"+user.Email, otp, 5*time.Minute)
	if err != nil {
		return err
	}

	emailBody := utils.BuildEmailBody(user.Name, otp)

	go func() {
		err := mail.SendEmail(user.Email, "Kode Verifikasi Keamanan - MediLink", emailBody)
		if err != nil {
			fmt.Printf("Gagal mengirim email ke %s: %v\n", user.Email, err)
		}
	}()

	return nil
}

func (u *userUsecase) VerifyOTP(ctx context.Context, userID uuid.UUID, inputOTP string) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return nil
	}

	key := "otp:" + user.Email

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

func (u *userUsecase) OnBoardPatient(ctx context.Context, userID uuid.UUID, data dto.PatientCreateRequestDTO) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if !user.IsVerified {
		return errs.ErrUserNotVerified
	}

	var patient entity.Patient
	patient.UserID = userID
	data.AssignToEntity(&patient)
	_, err = u.patientRepo.Create(ctx, &patient)
	if err != nil {
		return err
	}

	user.Role = enum.UserRole(enum.RolePatient)
	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
