package entity

import "github.com/google/uuid"

type Medicine struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `gorm:"type:varchar(255);not null"`

	Dosage *string `gorm:"type:varchar(100)"`
	Price  float64 `gorm:"type:numeric(12,2);not null"`
	Stock  int     `gorm:"default:0"`

	RequiresPrescription bool `gorm:"default:true"`
}
