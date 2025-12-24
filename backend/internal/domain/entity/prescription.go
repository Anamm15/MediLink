package entity

import (
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PatientID uuid.UUID  `gorm:"type:uuid;not null"`
	DoctorID  uuid.UUID  `gorm:"type:uuid;not null"`
	ClinicID  *uuid.UUID `gorm:"type:uuid"`

	Notes      *string `gorm:"type:text"`
	IsRedeemed bool    `gorm:"default:false"`

	CreatedAt time.Time              `gorm:"type:timestamptz;default:now()"`
	Patient   *Patient               `gorm:"foreignKey:PatientID"`
	Doctor    *Doctor                `gorm:"foreignKey:DoctorID"`
	Clinic    *Clinic                `gorm:"foreignKey:ClinicID"`
	Medicines []PrescriptionMedicine `gorm:"foreignKey:PrescriptionID"`
}
