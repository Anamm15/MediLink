package entity

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`

	Title   string `gorm:"type:varchar(255);not null"`
	Message string `gorm:"type:text;not null"`

	Type     *string `gorm:"type:varchar(100)"`
	Priority string  `gorm:"type:varchar(20);default:'normal'"`

	ReadAt *time.Time `gorm:"type:timestamptz"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
