package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type DoctorProfileResponseDTO struct {
	DoctorID     uuid.UUID `json:"doctor_id" gorm:"column:id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	ClinicName   *string   `json:"clinic_name"`
	ClinicAdress *string   `json:"clinic_address"`

	Specialization  string         `json:"specialization"`
	LicenseNumber   string         `json:"license_number"`
	Bio             *string        `json:"bio"`
	ExperienceYears int            `json:"experience_years"`
	Education       datatypes.JSON `json:"education"`

	RatingCount int     `json:"rating_count"`
	RatingTotal float64 `json:"rating_total"`
	ReviewCount int     `json:"review_count"`

	Schedule []DoctorScheduleResponseDTO `json:"schedule"`
}

type DoctorUpdateRequestDTO struct {
	Specialization  *string         `json:"specialization"`
	LicenseNumber   *string         `json:"license_number"`
	ExperienceYears *int            `json:"experience_years"`
	Education       *datatypes.JSON `json:"education"`
	Bio             *string         `json:"bio"`
}

type DoctorScheduleResponseDTO struct {
	ID        uuid.UUID        `json:"id"`
	DayOfWeek enum.ScheduleDay `json:"day_of_week"`
	StartTime time.Time        `json:"start_time"`
	EndTime   time.Time        `json:"end_time"`

	IsActive bool `json:"is_active"`
	MaxQuota *int `json:"max_quota"`
}

type DoctorCreateScheduleRequestDTO struct {
	DoctorID  uuid.UUID        `json:"doctor_id" binding:"required" validate:"required"`
	DayOfWeek enum.ScheduleDay `json:"day_of_week" binding:"required" validate:"required"`
	StartTime time.Time        `json:"start_time" binding:"required" validate:"required"`
	EndTime   time.Time        `json:"end_time" binding:"required" validate:"required"`
	IsActive  bool             `json:"is_active"`
	MaxQuota  *int             `json:"max_quota"`
}

type DoctorUpdateScheduleRequestDTO struct {
	DayOfWeek *enum.ScheduleDay `json:"day_of_week"`
	StartTime *time.Time        `json:"start_time"`
	EndTime   *time.Time        `json:"end_time"`
	IsActive  *bool             `json:"is_active"`
	MaxQuota  *int              `json:"max_quota"`
}

func MapEntityToDoctorResponseDTO(entity *entity.Doctor) DoctorProfileResponseDTO {
	return DoctorProfileResponseDTO{
		DoctorID:     entity.ID,
		Name:         entity.User.Name,
		Email:        entity.User.Email,
		PhoneNumber:  entity.User.PhoneNumber,
		ClinicName:   &entity.Clinic.Name,
		ClinicAdress: &entity.Clinic.Address,

		Specialization:  entity.Specialization,
		LicenseNumber:   entity.LicenseNumber,
		Education:       entity.Education,
		Bio:             entity.Bio,
		ExperienceYears: entity.ExperienceYears,

		RatingCount: entity.RatingCount,
		RatingTotal: entity.RatingTotal,
		ReviewCount: entity.ReviewCount,

		Schedule: MapListEntityDoctorScheduleToResponseDTO(entity.DoctorSchedule),
	}
}

func MapCreateScheduleRequestToEntity(dto *DoctorCreateScheduleRequestDTO) entity.DoctorSchedule {
	return entity.DoctorSchedule{
		DayOfWeek: dto.DayOfWeek,
		StartTime: dto.StartTime,
		EndTime:   dto.EndTime,
		IsActive:  dto.IsActive,
		MaxQuota:  dto.MaxQuota,
	}
}

func MapEntityDoctorScheduleToResponseDTO(entity *entity.DoctorSchedule) DoctorScheduleResponseDTO {
	return DoctorScheduleResponseDTO{
		ID:        entity.ID,
		DayOfWeek: entity.DayOfWeek,
		StartTime: entity.StartTime,
		EndTime:   entity.EndTime,
		IsActive:  entity.IsActive,
		MaxQuota:  entity.MaxQuota,
	}
}

func MapListEntityDoctorScheduleToResponseDTO(entity []entity.DoctorSchedule) []DoctorScheduleResponseDTO {
	var result []DoctorScheduleResponseDTO
	for _, schedule := range entity {
		result = append(result, MapEntityDoctorScheduleToResponseDTO(&schedule))
	}
	return result
}

func (dto *DoctorUpdateRequestDTO) AssignToEntity(doctor *entity.Doctor) {
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

func (dto *DoctorCreateScheduleRequestDTO) AssignToEntity(doctor *entity.DoctorSchedule) {
	doctor.DayOfWeek = dto.DayOfWeek
	doctor.StartTime = dto.StartTime
	doctor.EndTime = dto.EndTime
	doctor.IsActive = dto.IsActive
	doctor.MaxQuota = dto.MaxQuota
}

func (dto *DoctorUpdateScheduleRequestDTO) AssignToEntity(doctor *entity.DoctorSchedule) {
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
