package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/utils"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ClinicResponseDTO struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Type       string    `json:"type"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`

	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`

	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`

	InsurancePartners datatypes.JSON `json:"insurance_partners"`
	Facilities        datatypes.JSON `json:"facilities"`
	OpeningTime       datatypes.JSON `json:"opening_time"`

	Status          string     `json:"status"`
	Accreditation   *string    `json:"accreditation"`
	EstablishedDate *time.Time `json:"established_date"`

	CreatedAt time.Time `json:"created_at"`
}

type ClinicCreateRequestDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Code        string `json:"code" validate:"required,min=3,max=50"`
	Type        string `json:"type" validate:"required,min=3,max=100"`
	Address     string `json:"address" validate:"required"`
	City        string `json:"city" validate:"required,min=3,max=100"`
	Province    string `json:"province" validate:"required,min=3,max=100"`
	PostalCode  string `json:"postal_code" validate:"required,min=3,max=10"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20"`
	Email       string `json:"email" validate:"required,email"`

	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`

	OpeningTime       datatypes.JSON `json:"opening_time" validate:"required"`
	Facilities        datatypes.JSON `json:"facilities" validate:"required"`
	InsurancePartners datatypes.JSON `json:"insurance_partners" validate:"required"`

	Accreditation   *string `json:"accreditation"`
	EstablishedDate *string `json:"established_date"`
}

type ClinicUpdateRequestDTO struct {
	Name       *string `json:"name"`
	Code       *string `json:"code"`
	Type       *string `json:"type"`
	Address    *string `json:"address"`
	City       *string `json:"city"`
	Province   *string `json:"province"`
	PostalCode *string `json:"postal_code"`

	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`

	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`

	InsurancePartners *datatypes.JSON `json:"insurance_partners"`
	Facilities        *datatypes.JSON `json:"facilities"`
	OpeningTime       *datatypes.JSON `json:"opening_time"`

	Status          *string    `json:"status"`
	Accreditation   *string    `json:"accreditation"`
	EstablishedDate *time.Time `json:"established_date"`
}

func ToClinicResponseDTO(clinic *entity.Clinic) ClinicResponseDTO {
	return ClinicResponseDTO{
		ID:         clinic.ID,
		Name:       clinic.Name,
		Code:       clinic.Code,
		Type:       clinic.Type,
		Address:    clinic.Address,
		City:       clinic.City,
		Province:   clinic.Province,
		PostalCode: clinic.PostalCode,

		Latitude:  clinic.Latitude,
		Longitude: clinic.Longitude,

		PhoneNumber: clinic.PhoneNumber,
		Email:       clinic.Email,

		InsurancePartners: clinic.InsurancePartners,
		Facilities:        clinic.Facilities,
		OpeningTime:       clinic.OpeningTime,

		Status:          clinic.Status,
		Accreditation:   clinic.Accreditation,
		EstablishedDate: clinic.EstablishedDate,

		CreatedAt: clinic.CreatedAt,
	}
}

func (dto *ClinicCreateRequestDTO) AssignToEntity(clinic *entity.Clinic) {
	clinic.Name = dto.Name
	clinic.Code = dto.Code
	clinic.Type = dto.Type
	clinic.Address = dto.Address
	clinic.City = dto.City
	clinic.Province = dto.Province
	clinic.PostalCode = dto.PostalCode
	clinic.PhoneNumber = dto.PhoneNumber
	clinic.Email = dto.Email
	clinic.Latitude = dto.Latitude
	clinic.Longitude = dto.Longitude
	clinic.OpeningTime = dto.OpeningTime
	clinic.Facilities = dto.Facilities
	clinic.InsurancePartners = dto.InsurancePartners
	clinic.Accreditation = dto.Accreditation
	clinic.EstablishedDate = utils.ConvertStringToTime(*dto.EstablishedDate)
}

func (dto *ClinicUpdateRequestDTO) AssignToEntity(clinic *entity.Clinic) {
	if dto.Name != nil {
		clinic.Name = *dto.Name
	}
	if dto.Code != nil {
		clinic.Code = *dto.Code
	}
	if dto.Type != nil {
		clinic.Type = *dto.Type
	}
	if dto.Address != nil {
		clinic.Address = *dto.Address
	}
	if dto.City != nil {
		clinic.City = *dto.City
	}
	if dto.Province != nil {
		clinic.Province = *dto.Province
	}
	if dto.PostalCode != nil {
		clinic.PostalCode = *dto.PostalCode
	}

	if dto.Latitude != nil {
		clinic.Latitude = dto.Latitude
	}
	if dto.Longitude != nil {
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
	if dto.Status != nil {
		clinic.Status = *dto.Status
	}
	if dto.Accreditation != nil {
		clinic.Accreditation = dto.Accreditation
	}
	if dto.EstablishedDate != nil {
		clinic.EstablishedDate = dto.EstablishedDate
	}
}
