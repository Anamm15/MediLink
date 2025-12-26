package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecordRepository struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) repository.MedicalRecordRepository {
	return &MedicalRecordRepository{db: db}
}

func (r *MedicalRecordRepository) GetByPatientID(ctx context.Context, patientID uuid.UUID) ([]entity.MedicalRecord, error) {
	var medicalRecords []entity.MedicalRecord
	if err := r.db.WithContext(ctx).
		Preload("Patient", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Where("patient_id = ?", patientID).
		Find(&medicalRecords).Error; err != nil {
		return nil, err
	}
	return medicalRecords, nil
}

func (r *MedicalRecordRepository) GetByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]entity.MedicalRecord, error) {
	var medicalRecords []entity.MedicalRecord
	if err := r.db.WithContext(ctx).
		Preload("Patient", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Where("doctor_id = ?", doctorID).
		Find(&medicalRecords).Error; err != nil {
		return nil, err
	}
	return medicalRecords, nil
}

func (r *MedicalRecordRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.MedicalRecord, error) {
	var medicalRecord *entity.MedicalRecord
	if err := r.db.WithContext(ctx).
		Preload("Patient", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Where("id = ?", id).
		First(&medicalRecord).Error; err != nil {
		return nil, err
	}

	return medicalRecord, nil
}

func (r *MedicalRecordRepository) Create(ctx context.Context, medicalRecord *entity.MedicalRecord) error {
	if err := r.db.WithContext(ctx).
		Create(medicalRecord).Error; err != nil {
		return err
	}
	return nil
}

func (r *MedicalRecordRepository) Update(ctx context.Context, medicalRecord *entity.MedicalRecord) error {
	if err := r.db.WithContext(ctx).Save(medicalRecord).Error; err != nil {
		return err
	}
	return nil
}

func (r *MedicalRecordRepository) Delete(ctx context.Context, id uuid.UUID, doctorID uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Delete(&entity.MedicalRecord{}, "id = ? AND doctor_id = ?", id, doctorID).Error; err != nil {
		return err
	}
	return nil
}
