package dto

import (
	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type DoctorMinimumResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Specialization string    `json:"specialization"`
	PhoneNumber    string    `json:"phone_number"`
}

type DoctorClinicResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	City     string    `json:"city"`
	IsActive bool      `json:"is_active"`
}

type DoctorProfileResponse struct {
	ID          uuid.UUID `json:"id" gorm:"column:id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`

	Specialization  string         `json:"specialization"`
	LicenseNumber   string         `json:"license_number"`
	Bio             *string        `json:"bio"`
	ExperienceYears int            `json:"experience_years"`
	Education       datatypes.JSON `json:"education"`

	RatingCount int     `json:"rating_count"`
	RatingTotal float64 `json:"rating_total"`
	ReviewCount int     `json:"review_count"`

	Clinic []DoctorClinicResponse `json:"clinic"`
}

type DoctorUpdateRequest struct {
	Specialization  *string         `json:"specialization"`
	LicenseNumber   *string         `json:"license_number"`
	ExperienceYears *int            `json:"experience_years"`
	Education       *datatypes.JSON `json:"education"`
	Bio             *string         `json:"bio"`
}

type DoctorScheduleResponse struct {
	ID              uuid.UUID            `json:"id"`
	DayOfWeek       enum.ScheduleDay     `json:"day_of_week"`
	StartTime       string               `json:"start_time"`
	EndTime         string               `json:"end_time"`
	ConsultationFee float64              `json:"consultation_fee"`
	Type            enum.AppointmentType `json:"type"`

	IsActive bool `json:"is_active"`
	MaxQuota *int `json:"max_quota"`
}

type DoctorCreateScheduleRequest struct {
	ClinicID        *uuid.UUID           `json:"clinic_id"`
	DayOfWeek       enum.ScheduleDay     `json:"day_of_week" binding:"required" validate:"required"`
	StartTime       string               `json:"start_time" binding:"required" validate:"required"`
	EndTime         string               `json:"end_time" binding:"required" validate:"required"`
	ConsultationFee float64              `json:"consultation_fee" binding:"required" validate:"required"`
	Type            enum.AppointmentType `json:"type" binding:"required" validate:"required"`
	IsActive        bool                 `json:"is_active"`
	MaxQuota        *int                 `json:"max_quota"`
}

type DoctorUpdateScheduleRequest struct {
	DayOfWeek       *enum.ScheduleDay     `json:"day_of_week"`
	StartTime       *string               `json:"start_time"`
	EndTime         *string               `json:"end_time"`
	IsActive        *bool                 `json:"is_active"`
	ConsultationFee *float64              `json:"consultation_fee"`
	Type            *enum.AppointmentType `json:"type"`
	MaxQuota        *int                  `json:"max_quota"`
}

type DoctorUpdateStatusScheduleRequest struct {
	IsActive *bool `json:"is_active"`
}

func ToDoctorResponse(entity *entity.Doctor) DoctorProfileResponse {
	clinic := make([]DoctorClinicResponse, 0)
	for _, clinicEntity := range entity.DoctorClinicPlacements {
		clinic = append(clinic, DoctorClinicResponse{
			ID:       clinicEntity.ClinicID,
			Name:     clinicEntity.Clinic.Name,
			Address:  clinicEntity.Clinic.Address,
			City:     clinicEntity.Clinic.City,
			IsActive: clinicEntity.IsActive,
		})
	}

	return DoctorProfileResponse{
		ID:          entity.ID,
		Name:        entity.User.Name,
		Email:       entity.User.Email,
		PhoneNumber: entity.User.PhoneNumber,

		Specialization:  entity.Specialization,
		LicenseNumber:   entity.LicenseNumber,
		Education:       entity.Education,
		Bio:             entity.Bio,
		ExperienceYears: entity.ExperienceYears,

		RatingCount: entity.RatingCount,
		RatingTotal: entity.RatingTotal,
		ReviewCount: entity.ReviewCount,
		Clinic:      clinic,
	}
}

func ToDoctorScheduleResponse(entity *entity.DoctorSchedule) DoctorScheduleResponse {
	return DoctorScheduleResponse{
		ID:              entity.ID,
		DayOfWeek:       entity.DayOfWeek,
		StartTime:       entity.StartTime,
		EndTime:         entity.EndTime,
		IsActive:        entity.IsActive,
		ConsultationFee: entity.ConsultationFee,
		MaxQuota:        entity.MaxQuota,
		Type:            entity.Type,
	}
}

func ToListDoctorScheduleResponse(entity []entity.DoctorSchedule) []DoctorScheduleResponse {
	var result []DoctorScheduleResponse
	for _, schedule := range entity {
		result = append(result, ToDoctorScheduleResponse(&schedule))
	}
	return result
}

func (dto *DoctorUpdateRequest) ToModel(doctor *entity.Doctor) {
	if dto.Specialization != nil {
		doctor.Specialization = *dto.Specialization
	}
	if dto.LicenseNumber != nil {
		doctor.LicenseNumber = *dto.LicenseNumber
	}
	if dto.ExperienceYears != nil {
		doctor.ExperienceYears = *dto.ExperienceYears
	}
	if dto.Education != nil {
		doctor.Education = *dto.Education
	}
	if dto.Bio != nil {
		doctor.Bio = dto.Bio
	}
}

func (dto *DoctorCreateScheduleRequest) ToModel(doctor *entity.DoctorSchedule) {
	doctor.ClinicID = dto.ClinicID
	doctor.DayOfWeek = dto.DayOfWeek
	doctor.StartTime = dto.StartTime
	doctor.EndTime = dto.EndTime
	doctor.IsActive = dto.IsActive
	doctor.MaxQuota = dto.MaxQuota
	doctor.ConsultationFee = dto.ConsultationFee
	doctor.Type = dto.Type
}

func (dto *DoctorUpdateScheduleRequest) ToModel(doctor *entity.DoctorSchedule) {
	if dto.DayOfWeek != nil {
		doctor.DayOfWeek = *dto.DayOfWeek
	}
	if dto.StartTime != nil {
		doctor.StartTime = *dto.StartTime
	}
	if dto.EndTime != nil {
		doctor.EndTime = *dto.EndTime
	}
	if dto.IsActive != nil {
		doctor.IsActive = *dto.IsActive
	}
	if dto.ConsultationFee != nil {
		doctor.ConsultationFee = *dto.ConsultationFee
	}
	if dto.MaxQuota != nil {
		doctor.MaxQuota = dto.MaxQuota
	}
	if dto.Type != nil {
		doctor.Type = *dto.Type
	}
}
