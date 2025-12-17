package usecase

import (
	"context"

	"MediLink/internal/domain/entity"
	errs "MediLink/internal/domain/errors"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"
	"MediLink/internal/utils"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) Register(ctx context.Context, data dto.UserRegistrationRequestDTO) (dto.UserRegistrationResponseDTO, error) {
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return dto.UserRegistrationResponseDTO{}, err
	}

	user := &entity.User{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
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

func (u *userUsecase) GetProfile(ctx context.Context, userID uuid.UUID) (dto.UserResponseDTO, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return dto.MapUserToUserResponseDTO(user), nil
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

// func (u *userUsecase) OnBoardPatient(userID uuid.UUID, medicalHistory string) error {
// 	// return u.userRepo.OnBoardPatient(context.Background(), userID, medicalHistory)
// }
