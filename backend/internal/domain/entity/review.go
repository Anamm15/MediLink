package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	PatientID uuid.UUID  `gorm:"type:uuid;not null"`
	DoctorID  *uuid.UUID `gorm:"type:uuid"`
	ClinicID  *uuid.UUID `gorm:"type:uuid"`

	Rating  int     `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment *string `gorm:"type:text"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
