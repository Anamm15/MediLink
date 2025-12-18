package entity

import (
	"time"

	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName   string    `gorm:"type:varchar(255);not null"`
	LastName    string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"type:varchar(20);uniqueIndex;not null"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password    string    `gorm:"type:varchar(255);not null"`
	Address     *string   `gorm:"type:text"`

	Role   constants.UserRole   `gorm:"type:varchar(20);not null; default:'user'"`
	Status constants.UserStatus `gorm:"type:varchar(20);default:'active'"`

	BirthPlace *string           `gorm:"type:varchar(100)"`
	BirthDate  *time.Time        `gorm:"type:date"`
	Gender     *constants.Gender `gorm:"type:varchar(10);not null"`
	IsVerified bool              `gorm:"default:false"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
