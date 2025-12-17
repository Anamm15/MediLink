package entity

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	PatientID     uuid.UUID  `gorm:"type:uuid;not null"`
	DoctorID      uuid.UUID  `gorm:"type:uuid;not null"`
	AppointmentID *uuid.UUID `gorm:"type:uuid"`

	Diagnosis     *string `gorm:"type:text"`
	TreatmentPlan *string `gorm:"type:text"`

	NextAppointmentDate *time.Time `gorm:"type:date"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
