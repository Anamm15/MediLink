package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type clinicRepository struct {
	db *gorm.DB
}

func NewClinicRepository(db *gorm.DB) repository.ClinicRepository {
	return &clinicRepository{
		db: db,
	}
}

func (r *clinicRepository) Create(ctx context.Context, clinic *entity.Clinic) (*entity.Clinic, error) {
	if err := r.db.WithContext(ctx).
		Create(clinic).Error; err != nil {
		return nil, err
	}
	return clinic, nil
}

func (r *clinicRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Clinic, error) {
	var clinic entity.Clinic
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&clinic).Error; err != nil {
		return nil, err
	}
	return &clinic, nil
}

func (r *clinicRepository) GetAll(ctx context.Context, limit int, offset int) ([]entity.Clinic, error) {
	var clinics []entity.Clinic
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&clinics).Error; err != nil {
		return nil, err
	}
	return clinics, nil
}

func (r *clinicRepository) Find(ctx context.Context, name string, limit int, offset int) ([]entity.Clinic, error) {
	var clinics []entity.Clinic
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Where("name LIKE ?", "%"+name+"%").
		Find(&clinics).Error; err != nil {
		return nil, err
	}
	return clinics, nil
}

func (r *clinicRepository) Update(ctx context.Context, clinic *entity.Clinic) error {
	if err := r.db.WithContext(ctx).
		Model(clinic).
		Omit("id").
		Updates(clinic).Error; err != nil {
		return err
	}
	return nil
}

func (r *clinicRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&entity.Clinic{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
