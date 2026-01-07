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

func (r *DoctorClinicPlacementRepository) GetByID(tx *gorm.DB, scheduleID uuid.UUID) (entity.DoctorClinicPlacement, error) {
	var DoctorClinicPlacement entity.DoctorClinicPlacement
	if err := tx.First(&DoctorClinicPlacement, "id = ?", scheduleID).Error; err != nil {
		return entity.DoctorClinicPlacement{}, err
	}
	return DoctorClinicPlacement, nil
}

func (r *DoctorClinicPlacementRepository) Add(ctx context.Context, DoctorClinicPlacement *entity.DoctorClinicPlacement) error {
	if err := r.db.Create(DoctorClinicPlacement).Error; err != nil {
		return err
	}
	return nil
}

func (r *DoctorClinicPlacementRepository) Delete(ctx context.Context, doctorID uuid.UUID, clinicID uuid.UUID) error {
	if err := r.db.Delete(&entity.DoctorClinicPlacement{}, "doctor_id = ? AND clinic_id = ?", doctorID, clinicID).Error; err != nil {
		return err
	}
	return nil
}
