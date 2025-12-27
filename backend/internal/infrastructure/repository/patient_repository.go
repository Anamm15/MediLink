package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) repository.PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Patient, error) {
	var patient entity.Patient
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&patient).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (r *PatientRepository) Create(ctx context.Context, patient *entity.Patient) (*entity.Patient, error) {
	if err := r.db.WithContext(ctx).
		Create(patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func (r *PatientRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Patient, error) {
	var patient entity.Patient
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&patient).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (r *PatientRepository) Update(ctx context.Context, patient *entity.Patient) error {
	if err := r.db.WithContext(ctx).
		Model(patient).
		Omit("id", "user_id").
		Updates(patient).Error; err != nil {
		return err
	}
	return nil
}
