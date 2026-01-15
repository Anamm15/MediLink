package entity

import (
	"time"

	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string        `gorm:"type:varchar(255);uniqueIndex;not null;idx:idx_email"`
	Password  string        `gorm:"type:varchar(255);not null"`
	Role      enum.UserRole `gorm:"type:varchar(20);not null; default:'user'"`
	AvatarUrl string        `gorm:"type:text; default:null"`

	Name        string          `gorm:"type:varchar(255);not null"`
	PhoneNumber string          `gorm:"type:varchar(20);uniqueIndex;not null"`
	Status      enum.UserStatus `gorm:"type:varchar(20);default:'active'"`
	IsVerified  bool            `gorm:"default:false"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
