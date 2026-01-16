package repository

import (
	"context"
	"fmt"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PrescriptionRepository struct {
	db *gorm.DB
}

func NewPrescriptionRepository(db *gorm.DB) repository.PrescriptionRepository {
	return &PrescriptionRepository{
		db: db,
	}
}

func (r *PrescriptionRepository) GetByPatient(ctx context.Context, patientID uuid.UUID, limit int, offset int, isRedeemed bool) ([]entity.Prescription, int64, error) {
	var (
		prescriptions []entity.Prescription
		total         int64
	)

	baseQuery := r.db.WithContext(ctx).
		Model(&entity.Prescription{}).
		Where("patient_id = ? AND is_redeemed = ?", patientID, isRedeemed)

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := baseQuery.
		Preload("Doctor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id", "specialization")
		}).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "avatar_url")
		}).
		Preload("Medicines").
		Preload("Medicines.Medicine").
		Order("prescriptions.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

func (r *PrescriptionRepository) GetByDoctor(ctx context.Context, doctorID uuid.UUID, limit int, offset int) ([]entity.Prescription, int64, error) {
	var (
		prescriptions []entity.Prescription
		total         int64
	)

	baseQuery := r.db.WithContext(ctx).
		Model(&entity.Prescription{}).
		Where("doctor_id = ?", doctorID)

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := baseQuery.
		Preload("Patient", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "email", "phone_number", "avatar_url")
		}).
		Preload("Medicines").
		Preload("Medicines.Medicine").
		Order("prescriptions.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

func (r *PrescriptionRepository) GetDetailByID(ctx context.Context, id uuid.UUID) (*entity.Prescription, error) {
	var prescription entity.Prescription

	err := r.db.WithContext(ctx).
		Preload("Patient", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "avatar_url")
		}).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id", "specialization")
		}).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "avatar_url")
		}).
		Preload("Medicines").
		Preload("Medicines.Medicine").
		Find(&prescription, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &prescription, nil
}

func (r *PrescriptionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Prescription, error) {
	var prescription entity.Prescription

	err := r.db.WithContext(ctx).
		Find(&prescription, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &prescription, nil
}

func (r *PrescriptionRepository) Create(ctx context.Context, prescription *entity.Prescription) error {
	// Start a transaction
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Loop through medicines to Check Stock and Deduct
		for _, item := range prescription.Medicines {
			var medicine entity.Medicine

			// Lock the row to prevent race conditions
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&medicine, "id = ?", item.MedicineID).Error; err != nil {
				return fmt.Errorf("medicine not found: %w", err)
			}
		}

		// 2. Create the Prescription (and the items automatically)
		if err := tx.Create(prescription).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *PrescriptionRepository) Update(ctx context.Context, prescription *entity.Prescription) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.Prescription{}).
		Where("id = ?", prescription.ID).
		Omit("id", "Doctor", "Patient", "Medicines", "created_at").
		Updates(map[string]interface{}{
			"notes":       prescription.Notes,
			"is_redeemed": prescription.IsRedeemed,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PrescriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Delete(&entity.Prescription{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
