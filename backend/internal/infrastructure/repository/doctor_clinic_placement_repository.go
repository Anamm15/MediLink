package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorClinicPlacementRepository struct {
	db *gorm.DB
}

func NewDoctorClinicPlacementRepository(db *gorm.DB) repository.DoctorClinicPlacementRepository {
	return &DoctorClinicPlacementRepository{db: db}
}

func (dcpr *DoctorClinicPlacementRepository) Add(ctx context.Context, DoctorClinicPlacement *entity.DoctorClinicPlacement) error {
	if err := dcpr.db.Create(DoctorClinicPlacement).Error; err != nil {
		return err
	}
	return nil
}

func (dcpr *DoctorClinicPlacementRepository) Delete(ctx context.Context, doctorID uuid.UUID, clinicID uuid.UUID) error {
	if err := dcpr.db.Delete(&entity.DoctorClinicPlacement{}, "doctor_id = ? AND clinic_id = ?", doctorID, clinicID).Error; err != nil {
		return err
	}
	return nil
}
