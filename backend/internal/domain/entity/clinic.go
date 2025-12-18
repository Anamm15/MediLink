package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Clinic struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Code       string    `gorm:"type:varchar(50);uniqueIndex; not null"`
	Type       string    `gorm:"type:varchar(100); not null"`
	Address    string    `gorm:"type:text;not null"`
	City       string    `gorm:"type:varchar(100); not null"`
	Province   string    `gorm:"type:varchar(100); not null"`
	PostalCode string    `gorm:"type:varchar(10); not null"`

	Latitude  *float64 `gorm:"type:numeric(9,6)"`
	Longitude *float64 `gorm:"type:numeric(9,6)"`

	PhoneNumber string `gorm:"type:varchar(20); not null"`
	Email       string `gorm:"type:varchar(255); not null"`

	InsurancePartners datatypes.JSON `gorm:"type:jsonb"`
	Facilities        datatypes.JSON `gorm:"type:jsonb"`
	OpeningTime       datatypes.JSON `gorm:"type:jsonb"`

	Status          string     `gorm:"type:varchar(50);default:'active'"`
	Accreditation   *string    `gorm:"type:varchar(100)"`
	EstablishedDate *time.Time `gorm:"type:date"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
