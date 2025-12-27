package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) repository.AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) GetAll(ctx context.Context, limit int, offset int) ([]entity.Appointment, error) {
	var appointments []entity.Appointment
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id", "specialization") }).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name", "phone_number") }).
		Preload("Patient", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetByID(ctx context.Context, appointmentID uuid.UUID) (*entity.Appointment, error) {
	var appointment entity.Appointment
	if err := r.db.WithContext(ctx).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id", "specialization") }).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name", "phone_number") }).
		Preload("Patient", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		First(&appointment, "id = ?", appointmentID).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *AppointmentRepository) GetByDoctorID(ctx context.Context, doctorID uuid.UUID, limit int, offset int) ([]entity.Appointment, error) {
	var appointments []entity.Appointment
	if err := r.db.WithContext(ctx).
		Where("doctor_id = ?", doctorID).
		Limit(limit).
		Offset(offset).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id", "specialization") }).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name", "phone_number") }).
		Preload("Patient", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetByPatientID(ctx context.Context, patientID uuid.UUID, limit int, offset int) ([]entity.Appointment, error) {
	var appointments []entity.Appointment
	if err := r.db.WithContext(ctx).
		Where("patient_id = ?", patientID).
		Limit(limit).
		Offset(offset).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id", "specialization") }).
		Preload("Doctor.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name", "phone_number") }).
		Preload("Patient", func(db *gorm.DB) *gorm.DB { return db.Select("id", "user_id") }).
		Preload("Patient.User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) Create(ctx context.Context, appointment *entity.Appointment) error {
	if err := r.db.WithContext(ctx).Create(appointment).Error; err != nil {
		return err
	}
	return nil
}

func (r *AppointmentRepository) UpdateStatus(ctx context.Context, appointmentID uuid.UUID, status enum.AppointmentStatus) error {
	if err := r.db.WithContext(ctx).Model(&entity.Appointment{}).
		Where("id = ?", appointmentID).
		Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *AppointmentRepository) Delete(ctx context.Context, appointmentID uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Delete(&entity.Appointment{}, "id = ?", appointmentID).Error; err != nil {
		return err
	}
	return nil
}
