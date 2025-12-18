package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type DoctorProfileResponseDTO struct {
	DoctorID       uuid.UUID                   `json:"doctor_id" gorm:"column:id"`
	FirstName      string                      `json:"first_name"`
	LastName       string                      `json:"last_name"`
	Email          string                      `json:"email"`
	PhoneNumber    string                      `json:"phone_number"`
	ClinicName     *string                     `json:"clinic_name"`
	ClinicAdress   *string                     `json:"clinic_address"`
	Specialization string                      `json:"specialization"`
	LicenseNumber  string                      `json:"license_number"`
	Experience     datatypes.JSON              `json:"experience"`
	Education      datatypes.JSON              `json:"education"`
	IsActive       bool                        `json:"is_active"`
	Rating         float64                     `json:"rating"`
	TotalReviews   int                         `json:"total_reviews"`
	Bio            *string                     `json:"bio"`
	Schedule       []DoctorScheduleResponseDTO `json:"schedule"`
}

type DoctorScheduleResponseDTO struct {
	ID              uuid.UUID                 `json:"id"`
	DayOfWeek       constants.ScheduleDay     `json:"day_of_week"`
	StartTime       time.Time                 `json:"start_time"`
	EndTime         time.Time                 `json:"end_time"`
	Type            constants.AppointmentType `json:"type"`
	Location        *string                   `json:"location"`
	MaxAppointments *int                      `json:"max_appointments"`
}

type DoctorUpdateRequestDTO struct {
	Specialization *string         `json:"specialization"`
	LicenseNumber  *string         `json:"license_number"`
	Experience     *datatypes.JSON `json:"experience"`
	Education      *datatypes.JSON `json:"education"`
	IsActive       *bool           `json:"is_active"`
	Bio            *string         `json:"bio"`
}

type DoctorCreateScheduleRequestDTO struct {
	DoctorID        uuid.UUID                 `json:"doctor_id" binding:"required" validate:"required"`
	DayOfWeek       constants.ScheduleDay     `json:"day_of_week" binding:"required" validate:"required"`
	StartTime       time.Time                 `json:"start_time" binding:"required" validate:"required"`
	EndTime         time.Time                 `json:"end_time" binding:"required" validate:"required"`
	Type            constants.AppointmentType `json:"type" binding:"required" validate:"required"`
	Location        *string                   `json:"location" binding:"required" validate:"required"`
	MaxAppointments *int                      `json:"max_appointments" binding:"required" validate:"required"`
}

type DoctorUpdateScheduleRequestDTO struct {
	DayOfWeek       *constants.ScheduleDay     `json:"day_of_week"`
	StartTime       *time.Time                 `json:"start_time"`
	EndTime         *time.Time                 `json:"end_time"`
	Type            *constants.AppointmentType `json:"type"`
	Location        *string                    `json:"location"`
	MaxAppointments *int                       `json:"max_appointments"`
}

func MapEntityToDoctorResponseDTO(entity *entity.Doctor) DoctorProfileResponseDTO {
	return DoctorProfileResponseDTO{
		DoctorID:       entity.ID,
		FirstName:      entity.User.FirstName,
		LastName:       entity.User.LastName,
		Email:          entity.User.Email,
		PhoneNumber:    entity.User.PhoneNumber,
		ClinicName:     &entity.Clinic.Name,
		ClinicAdress:   &entity.Clinic.Address,
		Specialization: entity.Specialization,
		LicenseNumber:  entity.LicenseNumber,
		Experience:     entity.Experience,
		Education:      entity.Education,
		IsActive:       entity.IsActive,
		Rating:         entity.Rating,
		TotalReviews:   entity.TotalReviews,
		Bio:            entity.Bio,
		Schedule:       MapListEntityDoctorScheduleToResponseDTO(entity.DoctorSchedule),
	}
}

func MapCreateScheduleRequestToEntity(dto *DoctorCreateScheduleRequestDTO) entity.DoctorSchedule {
	return entity.DoctorSchedule{
		DayOfWeek:       dto.DayOfWeek,
		StartTime:       dto.StartTime,
		EndTime:         dto.EndTime,
		Type:            dto.Type,
		Location:        dto.Location,
		MaxAppointments: dto.MaxAppointments,
	}
}

func MapEntityDoctorScheduleToResponseDTO(entity *entity.DoctorSchedule) DoctorScheduleResponseDTO {
	return DoctorScheduleResponseDTO{
		ID:              entity.ID,
		DayOfWeek:       entity.DayOfWeek,
		StartTime:       entity.StartTime,
		EndTime:         entity.EndTime,
		Type:            entity.Type,
		Location:        entity.Location,
		MaxAppointments: entity.MaxAppointments,
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
	if dto.Experience != nil {
		doctor.Experience = *dto.Experience
	}
	if dto.Education != nil {
		doctor.Education = *dto.Education
	}
	if dto.IsActive != nil {
		doctor.IsActive = *dto.IsActive
	}
	if dto.Bio != nil {
		doctor.Bio = dto.Bio
	}
}

func (dto *DoctorCreateScheduleRequestDTO) AssignToEntity(doctor *entity.DoctorSchedule) {
	doctor.DayOfWeek = dto.DayOfWeek
	doctor.StartTime = dto.StartTime
	doctor.EndTime = dto.EndTime
	doctor.Type = dto.Type
	doctor.Location = dto.Location
	doctor.MaxAppointments = dto.MaxAppointments
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
	if dto.Type != nil {
		doctor.Type = *dto.Type
	}
	if dto.Location != nil {
		doctor.Location = dto.Location
	}
	if dto.MaxAppointments != nil {
		doctor.MaxAppointments = dto.MaxAppointments
	}
}
