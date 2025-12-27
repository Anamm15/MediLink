package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

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

	Clinic   []DoctorClinicResponse   `json:"clinic"`
	Schedule []DoctorScheduleResponse `json:"schedule"`
}

type DoctorUpdateRequest struct {
	Specialization  *string         `json:"specialization"`
	LicenseNumber   *string         `json:"license_number"`
	ExperienceYears *int            `json:"experience_years"`
	Education       *datatypes.JSON `json:"education"`
	Bio             *string         `json:"bio"`
}

type DoctorScheduleResponse struct {
	ID        uuid.UUID        `json:"id"`
	DayOfWeek enum.ScheduleDay `json:"day_of_week"`
	StartTime time.Time        `json:"start_time"`
	EndTime   time.Time        `json:"end_time"`

	IsActive bool `json:"is_active"`
	MaxQuota *int `json:"max_quota"`
}

type DoctorCreateScheduleRequest struct {
	DoctorID  uuid.UUID        `json:"doctor_id" binding:"required" validate:"required"`
	ClinicID  *uuid.UUID       `json:"clinic_id"`
	DayOfWeek enum.ScheduleDay `json:"day_of_week" binding:"required" validate:"required"`
	StartTime time.Time        `json:"start_time" binding:"required" validate:"required"`
	EndTime   time.Time        `json:"end_time" binding:"required" validate:"required"`
	IsActive  bool             `json:"is_active"`
	MaxQuota  *int             `json:"max_quota"`
}

type DoctorUpdateScheduleRequest struct {
	DayOfWeek *enum.ScheduleDay `json:"day_of_week"`
	StartTime *time.Time        `json:"start_time"`
	EndTime   *time.Time        `json:"end_time"`
	IsActive  *bool             `json:"is_active"`
	MaxQuota  *int              `json:"max_quota"`
}

func ToDoctorResponse(entity *entity.Doctor) DoctorProfileResponse {
	var clinic []DoctorClinicResponse
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

		Schedule: ToListDoctorScheduleResponse(entity.DoctorSchedules),
	}
}

func ToDoctorScheduleResponse(entity *entity.DoctorSchedule) DoctorScheduleResponse {
	return DoctorScheduleResponse{
		ID:        entity.ID,
		DayOfWeek: entity.DayOfWeek,
		StartTime: entity.StartTime,
		EndTime:   entity.EndTime,
		IsActive:  entity.IsActive,
		MaxQuota:  entity.MaxQuota,
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
	doctor.DoctorID = dto.DoctorID
	doctor.ClinicID = dto.ClinicID
	doctor.DayOfWeek = dto.DayOfWeek
	doctor.StartTime = dto.StartTime
	doctor.EndTime = dto.EndTime
	doctor.IsActive = dto.IsActive
	doctor.MaxQuota = dto.MaxQuota
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
	if dto.MaxQuota != nil {
		doctor.MaxQuota = dto.MaxQuota
	}
}
