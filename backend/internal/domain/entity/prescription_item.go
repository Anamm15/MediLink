package entity

import "github.com/google/uuid"

type PrescriptionItem struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PrescriptionID uuid.UUID `gorm:"type:uuid;not null"`
	MedicineID     uuid.UUID `gorm:"type:uuid;not null"`

	Quantity     int     `gorm:"not null"`
	Instructions *string `gorm:"type:text"`

	Medicine Medicine `gorm:"foreignKey:MedicineID"`
}
