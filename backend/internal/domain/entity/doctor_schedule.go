package entity

import (
	"time"

	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DoctorID uuid.UUID `gorm:"type:uuid;not null; index:idx_doctor_schedules_on_doctor_id"`

	DayOfWeek constants.ScheduleDay     `gorm:"type:varchar(20);not null;index:idx_doctor_schedules_on_doctor_id"`
	StartTime time.Time                 `gorm:"type:time;not null;index:idx_doctor_schedules_on_doctor_id"`
	EndTime   time.Time                 `gorm:"type:time;not null;index:idx_doctor_schedules_on_doctor_id"`
	Type      constants.AppointmentType `gorm:"type:varchar(20);not null"`

	Location        *string `gorm:"type:text"`
	MaxAppointments *int

	IsActive bool `gorm:"default:true"`
}
