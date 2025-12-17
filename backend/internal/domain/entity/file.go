package entity

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	FileName string `gorm:"type:varchar(255);not null"`
	URL      string `gorm:"type:text;not null"`

	FileableID   uuid.UUID `gorm:"type:uuid;not null; index:idx_files_fileable"`
	FileableType string    `gorm:"type:varchar(100);not null; index:idx_files_fileable"`

	MimeType  *string `gorm:"type:varchar(100)"`
	SizeBytes *int    `gorm:"type:int"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
