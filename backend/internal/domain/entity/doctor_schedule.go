package entity

import (
	"time"

	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DoctorID uuid.UUID `gorm:"type:uuid;not null"`

	DayOfWeek constants.ScheduleDay     `gorm:"type:schedule_day;not null"`
	StartTime time.Time                 `gorm:"type:time;not null"`
	EndTime   time.Time                 `gorm:"type:time;not null"`
	Type      constants.AppointmentType `gorm:"type:appointment_type;not null"`

	Location        *string `gorm:"type:text"`
	MaxAppointments *int

	IsActive bool `gorm:"default:true"`

	// Composite unique index
	// doctor_id, day_of_week, start_time, end_time
}
