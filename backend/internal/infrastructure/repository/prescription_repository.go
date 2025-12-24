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

func (pr *PrescriptionRepository) GetByPatient(ctx context.Context, userID uuid.UUID,
) ([]entity.Prescription, error) {
	var prescriptions []entity.Prescription

	err := pr.db.WithContext(ctx).
		Joins("JOIN patients ON patients.id = prescriptions.patient_id").
		Where("patients.user_id = ?", userID).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id", "specialization")
		}).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Medicines").
		Preload("Medicines.Medicine").
		Find(&prescriptions).Error
	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (pr *PrescriptionRepository) GetByDoctor(ctx context.Context, userID uuid.UUID) ([]entity.Prescription, error) {
	var prescriptions []entity.Prescription

	err := pr.db.WithContext(ctx).
		Joins("JOIN doctors ON doctors.id = prescriptions.doctor_id").
		Where("doctors.user_id = ?", userID).
		Preload("Patient", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "user_id")
		}).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Medicines").
		Preload("Medicines.Medicine").
		Find(&prescriptions).Error
	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (pr *PrescriptionRepository) GetDetailByID(ctx context.Context, id uuid.UUID) (*entity.Prescription, error) {
	var prescription entity.Prescription

	err := pr.db.WithContext(ctx).
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

func (pr *PrescriptionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Prescription, error) {
	var prescription entity.Prescription

	err := pr.db.WithContext(ctx).
		Find(&prescription, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &prescription, nil
}

func (pr *PrescriptionRepository) Create(ctx context.Context, prescription *entity.Prescription) error {
	// Start a transaction
	return pr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (pr *PrescriptionRepository) Update(ctx context.Context, prescription *entity.Prescription) error {
	if err := pr.db.WithContext(ctx).
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

func (pr *PrescriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := pr.db.WithContext(ctx).
		Delete(&entity.Prescription{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
