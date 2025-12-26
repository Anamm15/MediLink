package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) repository.DoctorRepository {
	return &doctorRepository{db: db}
}

func (r *doctorRepository) GetWithSchedule(ctx context.Context, id uuid.UUID) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.WithContext(ctx).
		Preload("DoctorSchedules").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "email", "phone_number")
		}).
		Preload("DoctorClinicPlacements.Clinic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "address", "city", "is_active")
		}).
		First(&doctor, id).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.WithContext(ctx).
		First(&doctor, id).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) Find(ctx context.Context, name string, limit int, offset int) ([]entity.Doctor, error) {
	var doctors []entity.Doctor
	if err := r.db.WithContext(ctx).
		Model(&entity.Doctor{}).
		Preload("DoctorSchedules").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "email", "phone_number")
		}).
		Preload("DoctorClinicPlacements.Clinic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "address", "city", "is_active")
		}).
		Joins("JOIN users u ON u.id = doctors.user_id").
		Where("u.name ILIKE ?", "%"+name+"%").
		Limit(limit).
		Offset(offset).
		Find(&doctors).Error; err != nil {
		return nil, err
	}

	return doctors, nil
}

func (r *doctorRepository) Create(ctx context.Context, doctor *entity.Doctor) (*entity.Doctor, error) {
	if err := r.db.WithContext(ctx).
		Create(doctor).Error; err != nil {
		return nil, err
	}
	return doctor, nil
}

func (r *doctorRepository) Update(ctx context.Context, doctor *entity.Doctor) error {
	if err := r.db.WithContext(ctx).
		Model(doctor).
		Omit("id", "user_id").
		Updates(doctor).Error; err != nil {
		return err
	}
	return nil
}

func (r *doctorRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&entity.Doctor{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
