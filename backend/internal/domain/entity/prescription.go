package entity

import (
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PatientID       uuid.UUID `gorm:"type:uuid;not null"`
	DoctorID        uuid.UUID `gorm:"type:uuid;not null"`
	MedicalRecordID uuid.UUID `gorm:"type:uuid;not null"`

	Notes      *string `gorm:"type:text"`
	IsRedeemed bool    `gorm:"default:false"`

	Patient   Patient            `gorm:"foreignKey:PatientID"`
	Doctor    Doctor             `gorm:"foreignKey:DoctorID"`
	Medicines []PrescriptionItem `gorm:"foreignKey:PrescriptionID"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
