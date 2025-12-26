package entity

import (
	"github.com/google/uuid"
)

type DoctorClinicPlacement struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ClinicID uuid.UUID `gorm:"type:uuid;not null;index:idx_doctor_placements_clinic"`
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index:idx_doctor_placements_doctor"`

	ConsultationFee float64 `gorm:"type:numeric(12,2);not null"`
	IsActive        bool    `gorm:"default:true"`

	Doctor Doctor `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Clinic Clinic `gorm:"foreignKey:ClinicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
