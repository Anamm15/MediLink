package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	ID    uuid.UUID     `json:"id"`
	Email string        `json:"email"`
	Role  enum.UserRole `json:"role"`

	Name        string          `json:"name"`
	PhoneNumber string          `json:"phone_number"`
	Status      enum.UserStatus `json:"status"`
	IsVerified  bool            `json:"is_verified"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserProfileResponseDTO struct {
	User    UserResponseDTO    `json:"user"`
	Patient PatientResponseDTO `json:"patient"`
}

type UserUpdateProfileRequestDTO struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number" binding:"e164"`
}

type VerifyUserRequest struct {
	OTP string `json:"otp" binding:"required"`
}

type OnBoardPatientRequestDTO struct {
	MedicalHistory string `json:"medical_history" binding:"required"`
}

func MapUserToUserResponseDTO(user *entity.User) UserResponseDTO {
	return UserResponseDTO{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Status:      user.Status,
		IsVerified:  user.IsVerified,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func MapUserResponseDTOToUser(dto UserResponseDTO) entity.User {
	return entity.User{
		ID:          dto.ID,
		Name:        dto.Name,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
		Status:      dto.Status,
		IsVerified:  dto.IsVerified,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

func (req *UserUpdateProfileRequestDTO) ToModel(user *entity.User) {
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		user.PhoneNumber = *req.PhoneNumber
	}
}
