package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Doctor struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID   uuid.UUID  `gorm:"type:uuid;uniqueIndex;not null"`
	ClinicID *uuid.UUID `gorm:"type:uuid"`

	Specialization  string         `gorm:"type:varchar(100);not null;index:idx_doctor_specialization"`
	LicenseNumber   string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Bio             *string        `gorm:"type:text"`
	ExperienceYears int            `gorm:"not null"`
	Education       datatypes.JSON `gorm:"type:jsonb"`

	RatingCount int     `gorm:"default:0"`
	RatingTotal float64 `gorm:"type:numeric(3,2);default:0"`
	ReviewCount int     `gorm:"default:0"`

	DoctorClinicPlacements []DoctorClinicPlacement `gorm:"foreignKey:DoctorID"`
	DoctorSchedules        []DoctorSchedule        `gorm:"foreignKey:DoctorID"`
	User                   User                    `gorm:"foreignKey:UserID"`
	Appointments           []Appointment           `gorm:"foreignKey:DoctorID"`
	MedicalRecords         []MedicalRecord         `gorm:"foreignKey:DoctorID"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
