package entity

import (
	"time"

	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type Appointment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_patient"`
	DoctorID  uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_doctor_date"`
	ClinicID  uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_on_clinic_id"`

	AppointmentDate time.Time `gorm:"type:date;not null"`
	StartTime       time.Time `gorm:"type:timestamptz;not null;index:idx_appointments_doctor_date"`
	EndTime         time.Time `gorm:"type:timestamptz;not null"`

	Status enum.AppointmentStatus `gorm:"type:varchar(20);default:'pending';not null;index:idx_appointments_status"`
	Type   enum.AppointmentType   `gorm:"type:varchar(20);not null"`

	ConsultationFeeSnapshot float64 `gorm:"type:numeric(12,2);not null"`

	QueueNumber *int    `gorm:"type:int;"`
	MeetingLink *string `gorm:"type:text;"`

	SymptomComplaint *string `gorm:"type:text;"`
	DoctorNotes      *string `gorm:"type:text;"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
