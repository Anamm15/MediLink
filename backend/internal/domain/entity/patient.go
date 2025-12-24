package entity

import "github.com/google/uuid"

type Patient struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID            uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
	IdentityNumber    string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	BloodType         *string   `gorm:"type:varchar(5)"`
	WeightKg          *float64  `gorm:"type:numeric(5,2)"`
	HeightCm          *float64  `gorm:"type:numeric(5,2)"`
	Allergies         *string   `gorm:"type:text"`
	ChronicDiseases   *string   `gorm:"type:text"`
	EmergencyContact  *string   `gorm:"type:text"`
	InsuranceProvider *string   `gorm:"type:varchar(100)"`
	InsuranceNumber   *string   `gorm:"type:varchar(100)"`
	Occupation        *string   `gorm:"type:varchar(100)"`

	User User `gorm:"foreignKey:UserID"`
}
