package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/utils"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ClinicResponseDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	City    string    `json:"city"`

	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	PhoneNumber string   `json:"phone_number"`
	Email       string   `json:"email"`

	Facilities        datatypes.JSON  `json:"facilities"`
	OpeningTime       datatypes.JSON  `json:"opening_time"`
	InsurancePartners *datatypes.JSON `json:"insurance_partners"`

	IsActive      bool    `json:"is_active"`
	Accreditation *string `json:"accreditation"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClinicCreateRequestDTO struct {
	Name    string `json:"name" validate:"required,min=3,max=255"`
	Address string `json:"address" validate:"required"`
	City    string `json:"city" validate:"required,min=3,max=100"`

	PhoneNumber string   `json:"phone_number" validate:"required,min=3,max=20"`
	Email       string   `json:"email" validate:"required,email"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`

	OpeningTime       datatypes.JSON  `json:"opening_time" validate:"required"`
	Facilities        datatypes.JSON  `json:"facilities" validate:"required"`
	InsurancePartners *datatypes.JSON `json:"insurance_partners"`

	Accreditation *string `json:"accreditation"`
	IsActive      *bool   `json:"is_active"`
}

type ClinicUpdateRequestDTO struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	City    *string `json:"city"`

	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	PhoneNumber *string  `json:"phone_number"`
	Email       *string  `json:"email"`

	InsurancePartners *datatypes.JSON `json:"insurance_partners"`
	Facilities        *datatypes.JSON `json:"facilities"`
	OpeningTime       *datatypes.JSON `json:"opening_time"`

	IsActive      *bool   `json:"is_active"`
	Accreditation *string `json:"accreditation"`
}

type AssignDoctorRequest struct {
	ClinicID        uuid.UUID `json:"clinic_id" validate:"required"`
	DoctorID        uuid.UUID `json:"doctor_id" validate:"required"`
	ConsultationFee float64   `json:"consultation_fee" validate:"required"`
	IsActive        *bool     `json:"is_active"`
}

type RemoveDoctorRequest struct {
	ClinicID uuid.UUID `json:"clinic_id" validate:"required"`
	DoctorID uuid.UUID `json:"doctor_id" validate:"required"`
}

func ToClinicResponseDTO(clinic *entity.Clinic) ClinicResponseDTO {
	return ClinicResponseDTO{
		ID:      clinic.ID,
		Name:    clinic.Name,
		Address: clinic.Address,
		City:    clinic.City,

		Latitude:    clinic.Latitude,
		Longitude:   clinic.Longitude,
		PhoneNumber: clinic.PhoneNumber,
		Email:       clinic.Email,

		InsurancePartners: &clinic.InsurancePartners,
		Facilities:        clinic.Facilities,
		OpeningTime:       clinic.OpeningTime,

		IsActive:      clinic.IsActive,
		Accreditation: clinic.Accreditation,

		CreatedAt: clinic.CreatedAt,
		UpdatedAt: clinic.UpdatedAt,
	}
}

func (dto *ClinicCreateRequestDTO) AssignToEntity(clinic *entity.Clinic) {
	clinic.Name = dto.Name
	clinic.Address = dto.Address
	clinic.City = dto.City

	clinic.Latitude = dto.Latitude
	clinic.Longitude = dto.Longitude
	clinic.PhoneNumber = dto.PhoneNumber
	clinic.Email = dto.Email

	clinic.OpeningTime = dto.OpeningTime
	clinic.Facilities = dto.Facilities
	clinic.InsurancePartners = *dto.InsurancePartners

	clinic.Accreditation = dto.Accreditation
	clinic.IsActive = utils.GetBoolOrDefault(dto.IsActive, true)
}

func (dto *ClinicUpdateRequestDTO) AssignToEntity(clinic *entity.Clinic) {
	if dto.Name != nil {
		clinic.Name = *dto.Name
	}
	if dto.Address != nil {
		clinic.Address = *dto.Address
	}
	if dto.City != nil {
		clinic.City = *dto.City
	}

	if dto.Latitude != nil && dto.Longitude != nil {
		clinic.Latitude = dto.Latitude
		clinic.Longitude = dto.Longitude
	}
	if dto.PhoneNumber != nil {
		clinic.PhoneNumber = *dto.PhoneNumber
	}
	if dto.Email != nil {
		clinic.Email = *dto.Email
	}

	if dto.InsurancePartners != nil {
		clinic.InsurancePartners = *dto.InsurancePartners
	}
	if dto.Facilities != nil {
		clinic.Facilities = *dto.Facilities
	}
	if dto.OpeningTime != nil {
		clinic.OpeningTime = *dto.OpeningTime
	}

	if dto.IsActive != nil {
		clinic.IsActive = *dto.IsActive
	}
	if dto.Accreditation != nil {
		clinic.Accreditation = dto.Accreditation
	}
}

func (dto *AssignDoctorRequest) ToModel(doctor *entity.DoctorClinicPlacement) {
	doctor.ClinicID = dto.ClinicID
	doctor.DoctorID = dto.DoctorID
	doctor.ConsultationFee = dto.ConsultationFee
	doctor.IsActive = utils.GetBoolOrDefault(dto.IsActive, true)
}
