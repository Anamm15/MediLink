package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) repository.MedicineRepository {
	return &MedicineRepository{db: db}
}

func (r *MedicineRepository) GetAll(ctx context.Context, limit int, offset int) ([]entity.Medicine, int64, error) {
	var (
		medicines []entity.Medicine
		total     int64
	)

	baseQuery := r.db.WithContext(ctx).
		Model(entity.Medicine{})

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&medicines).Error; err != nil {
		return nil, 0, err
	}

	return medicines, total, nil
}

func (r *MedicineRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Medicine, error) {
	var medicine entity.Medicine
	if err := r.db.WithContext(ctx).
		First(&medicine, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *MedicineRepository) Search(ctx context.Context, name string, limit int, offset int) ([]entity.Medicine, int64, error) {
	var (
		medicines []entity.Medicine
		total     int64
	)

	baseQuery := r.db.WithContext(ctx).
		Model(entity.Medicine{}).
		Where(
			"name ILIKE ? OR generic_name ILIKE ?",
			"%"+name+"%",
			"%"+name+"%",
		)

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.
		Limit(limit).
		Offset(offset).
		Find(&medicines).Error; err != nil {
		return nil, 0, err
	}
	return medicines, total, nil
}

func (r *MedicineRepository) Create(ctx context.Context, medicine *entity.Medicine) error {
	if err := r.db.WithContext(ctx).
		Create(medicine).Error; err != nil {
		return err
	}
	return nil
}

func (r *MedicineRepository) Update(ctx context.Context, medicine *entity.Medicine) error {
	if err := r.db.WithContext(ctx).
		Model(medicine).
		Omit("id").
		Updates(medicine).Error; err != nil {
		return err
	}
	return nil
}

func (r *MedicineRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Delete(&entity.Medicine{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
