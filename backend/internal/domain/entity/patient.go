package entity

import (
	"time"

	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type Patient struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	IdentityNumber string      `gorm:"type:varchar(50);uniqueIndex;not null"`
	BirthDate      time.Time   `gorm:"type:date;not null"`
	Gender         enum.Gender `gorm:"type:varchar(10);not null"`
	BloodType      string      `gorm:"type:varchar(5);not null"`

	WeightKg               float64 `gorm:"type:numeric(5,2);not null"`
	HeightCm               float64 `gorm:"type:numeric(5,2);not null"`
	Allergies              *string `gorm:"type:text"`
	HistoryChronicDiseases *string `gorm:"type:text"`
	EmergencyContact       *string `gorm:"type:text"`

	InsuranceProvider *string `gorm:"type:varchar(100)"`
	InsuranceNumber   *string `gorm:"type:varchar(100)"`
	Occupation        *string `gorm:"type:varchar(100)"`

	User           User            `gorm:"foreignKey:UserID"`
	Appointments   []Appointment   `gorm:"foreignKey:PatientID"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
