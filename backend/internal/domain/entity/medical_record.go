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

	Subjective *string `gorm:"type:text"` // Patient's reported symptoms and complaints
	Objective  *string `gorm:"type:text"` // Clinician's observations and findings
	Assessment *string `gorm:"type:text"` // Diagnosis or clinical impression
	Plan       *string `gorm:"type:text"` // Treatment plan, prescriptions, follow-ups

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
