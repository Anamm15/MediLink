package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Clinic struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name    string    `gorm:"type:varchar(255);not null"`
	Address string    `gorm:"type:text;not null"`
	City    string    `gorm:"type:varchar(100);not null"`

	Latitude    *float64 `gorm:"type:numeric(9,6)"`
	Longitude   *float64 `gorm:"type:numeric(9,6)"`
	PhoneNumber string   `gorm:"type:varchar(20);not null"`
	Email       string   `gorm:"type:varchar(255);not null"`

	OpeningTime       datatypes.JSON `gorm:"type:jsonb;not null"`
	Facilities        datatypes.JSON `gorm:"type:jsonb;not null"`
	InsurancePartners datatypes.JSON `gorm:"type:jsonb"`

	IsActive      bool    `gorm:"default:true"`
	Accreditation *string `gorm:"type:varchar(100)"`

	Doctors      []Doctor      `gorm:"foreignKey:ClinicID"`
	Appointments []Appointment `gorm:"foreignKey:ClinicID"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
