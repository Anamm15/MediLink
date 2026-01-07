package entity

import (
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DoctorID uuid.UUID  `gorm:"type:uuid;not null; index:idx_doctor_schedules_on_doctor_id"`
	ClinicID *uuid.UUID `gorm:"type:uuid; index:idx_doctor_schedules_on_doctor_id"`

	DayOfWeek       enum.ScheduleDay `gorm:"type:varchar(20);not null;index:idx_doctor_schedules_on_doctor_id"`
	StartTime       string           `gorm:"type:time;not null;index:idx_doctor_schedules_on_doctor_id"`
	EndTime         string           `gorm:"type:time;not null;index:idx_doctor_schedules_on_doctor_id"`
	ConsultationFee float64          `gorm:"type:numeric(12,2);not null"`

	IsActive bool `gorm:"type:boolean;default:true"`
	MaxQuota *int `gorm:"type:int;default:null"`

	Doctor Doctor `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Clinic Clinic `gorm:"foreignKey:ClinicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
