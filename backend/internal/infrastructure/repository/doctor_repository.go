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

func (r *doctorRepository) GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.WithContext(ctx).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "email", "phone_number", "avatar_url")
		}).
		Preload("DoctorClinicPlacements.Clinic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "address", "city", "is_active")
		}).
		Where("user_id = ?", userID).
		First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) GetProfileByID(ctx context.Context, id uuid.UUID) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.WithContext(ctx).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "email", "phone_number", "avatar_url")
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

func (r *doctorRepository) Find(ctx context.Context, name string, limit int, offset int) ([]entity.Doctor, int64, error) {
	var (
		doctors []entity.Doctor
		total   int64
	)

	baseQuery := r.db.WithContext(ctx).
		Model(&entity.Doctor{}).
		Joins("JOIN users u ON u.id = doctors.user_id").
		Where("u.name ILIKE ?", "%"+name+"%")

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.
		Select(
			"doctors.id",
			"doctors.user_id",
			"doctors.specialization",
			"doctors.rating_total",
			"doctors.review_count",
		).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "email", "phone_number", "avatar_url")
		}).
		Limit(limit).
		Offset(offset).
		Find(&doctors).Error; err != nil {
		return nil, 0, err
	}

	return doctors, total, nil
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
