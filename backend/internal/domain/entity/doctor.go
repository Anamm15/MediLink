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

	Specialization  string  `gorm:"type:varchar(100);not null"`
	LicenseNumber   string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	ConsultationFee float64 `gorm:"type:numeric(12,2);default:0"`

	Experience datatypes.JSON `gorm:"type:jsonb"`
	Education  datatypes.JSON `gorm:"type:jsonb"`

	IsActive                 bool    `gorm:"default:true"`
	Rating                   float64 `gorm:"type:numeric(3,2);default:0"`
	TotalReviews             int     `gorm:"default:0"`
	AvailableForTelemedicine bool    `gorm:"default:false"`

	Bio *string `gorm:"type:text"`

	DoctorSchedule []DoctorSchedule `gorm:"foreignKey:DoctorID"`
	User           User             `gorm:"foreignKey:UserID"`
	Clinic         *Clinic          `gorm:"foreignKey:ClinicID"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
