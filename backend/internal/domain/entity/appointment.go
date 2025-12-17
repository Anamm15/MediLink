package entity

import (
	"time"

	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type Appointment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_on_patient_id"`
	DoctorID  uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_on_doctor_id"`

	ScheduleStartTime time.Time `gorm:"type:timestamptz;not null"`
	ScheduleEndTime   time.Time `gorm:"type:timestamptz;not null"`
	DurationMinutes   *int

	Status constants.AppointmentStatus `gorm:"type:appointment_status;default:'pending';not null"`
	Type   constants.AppointmentType   `gorm:"type:appointment_type;not null"`

	CanceledReason *string `gorm:"type:text"`
	Complaint      *string `gorm:"type:text"`
	Location       *string `gorm:"type:text"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
