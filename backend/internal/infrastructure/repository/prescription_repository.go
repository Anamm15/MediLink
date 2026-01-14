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

func (r *PrescriptionRepository) GetByPatient(ctx context.Context, patientID uuid.UUID) ([]entity.Prescription, error) {
	var prescriptions []entity.Prescription

	err := r.db.WithContext(ctx).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id", "specialization")
		}).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Medicines").
		Preload("Medicines.Medicine").
		Where("patient_id = ?", patientID).
		Find(&prescriptions).Error
	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (r *PrescriptionRepository) GetByDoctor(ctx context.Context, doctorID uuid.UUID) ([]entity.Prescription, error) {
	var prescriptions []entity.Prescription

	err := r.db.WithContext(ctx).
		Preload("Patient", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "email", "phone_number")
		}).
		Preload("Medicines").
		Preload("Medicines.Medicine").
		Where("doctor_id = ?", doctorID).
		Find(&prescriptions).Error
	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (r *PrescriptionRepository) GetDetailByID(ctx context.Context, id uuid.UUID) (*entity.Prescription, error) {
	var prescription entity.Prescription

	err := r.db.WithContext(ctx).
		Preload("Patient", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id", "specialization")
		}).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
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
