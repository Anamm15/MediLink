package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/constants"
	"MediLink/internal/utils"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	ID          uuid.UUID            `json:"id"`
	FirstName   string               `json:"first_name"`
	LastName    string               `json:"last_name"`
	PhoneNumber string               `json:"phone_number"`
	Email       string               `json:"email"`
	Address     *string              `json:"address"`
	Status      constants.UserStatus `json:"status"`
	BirthDate   string               `json:"birthdate"`
	BirthPlace  *string              `json:"birthplace"`
	Gender      *constants.Gender    `json:"gender"`
	IsVerified  bool                 `json:"is_verified"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type UserProfileResponseDTO struct {
	User    UserResponseDTO    `json:"user"`
	Patient PatientResponseDTO `json:"patient"`
}

type UserRegistrationRequestDTO struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	Gender      string `json:"gender" binding:"omitempty,oneof=male female other"`
	Address     string `json:"address" binding:"required"`
	BirthPlace  string `json:"birth_place" binding:"required"`
	BirthDate   string `json:"birth_date" binding:"required"`
}

type UserRegistrationResponseDTO struct {
	ID uuid.UUID `json:"id"`
	UserRegistrationRequestDTO
}

type UserLoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateProfileRequestDTO struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number" binding:"e164"`
	Address     *string `json:"address"`
	BirthPlace  *string `json:"birth_place"`
	BirthDate   *string `json:"birth_date"`
}

type UserChangePasswordRequestDTO struct {
	OldPassword *string `json:"old_password" binding:"required"`
	NewPassword *string `json:"new_password" binding:"required,min=8"`
}

type UserVerifyOTPRequestDTO struct {
	OTP string `json:"otp" binding:"required"`
}

type OnBoardPatientRequestDTO struct {
	MedicalHistory string `json:"medical_history" binding:"required"`
}

func MapUserToUserResponseDTO(user *entity.User) UserResponseDTO {
	return UserResponseDTO{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Address:     user.Address,
		Status:      user.Status,
		BirthDate:   utils.ConvertTimeToString(*user.BirthDate),
		BirthPlace:  user.BirthPlace,
		Gender:      user.Gender,
		IsVerified:  user.IsVerified,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func MapUserResponseDTOToUser(dto UserResponseDTO) entity.User {
	return entity.User{
		ID:          dto.ID,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
		Address:     dto.Address,
		Status:      dto.Status,
		BirthDate:   utils.ConvertStringToTime(dto.BirthDate),
		BirthPlace:  dto.BirthPlace,
		Gender:      dto.Gender,
		IsVerified:  dto.IsVerified,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

func (req *UserUpdateProfileRequestDTO) AssignToEntity(user *entity.User) {
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		user.PhoneNumber = *req.PhoneNumber
	}
	if req.Address != nil {
		user.Address = req.Address
	}
	if req.BirthPlace != nil {
		user.BirthPlace = req.BirthPlace
	}
	if req.BirthDate != nil {
		user.BirthDate = utils.ConvertStringToTime(*req.BirthDate)
	}
}
