package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorClinicPlacementRepository interface {
	Add(ctx context.Context, DoctorClinicPlacement *entity.DoctorClinicPlacement) error
	Delete(ctx context.Context, doctorID uuid.UUID, clinicID uuid.UUID) error
	GetByID(tx *gorm.DB, scheduleID uuid.UUID) (entity.DoctorClinicPlacement, error)
}
