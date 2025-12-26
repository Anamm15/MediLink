package repository

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClinicInventoryRepository struct {
	db *gorm.DB
}

func NewClinicInventoryRepository(db *gorm.DB) repository.ClinicInventoryRepository {
	return &ClinicInventoryRepository{db: db}
}

func (cir *ClinicInventoryRepository) GetAll(ctx context.Context) ([]entity.ClinicInventory, error) {
	var clinicInventories []entity.ClinicInventory
	if err := cir.db.
		WithContext(ctx).
		Preload("Medicine", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Clinic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Find(&clinicInventories).Error; err != nil {
		return nil, err
	}
	return clinicInventories, nil
}

func (cir *ClinicInventoryRepository) GetByClinicID(ctx context.Context, clinicID uuid.UUID) ([]entity.ClinicInventory, error) {
	var clinicInventories []entity.ClinicInventory
	if err := cir.db.
		WithContext(ctx).
		Preload("Medicine", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Clinic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Where("clinic_id = ?", clinicID).
		Find(&clinicInventories).Error; err != nil {
		return nil, err
	}
	return clinicInventories, nil
}

func (cir *ClinicInventoryRepository) GetByID(ctx context.Context, id uuid.UUID) (entity.ClinicInventory, error) {
	var clinicInventories entity.ClinicInventory
	if err := cir.db.
		WithContext(ctx).
		Preload("Medicine", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Clinic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Where("id = ?", id).
		Find(&clinicInventories).Error; err != nil {
		return entity.ClinicInventory{}, err
	}
	return clinicInventories, nil
}

func (cir *ClinicInventoryRepository) Create(ctx context.Context, clinicInventory *entity.ClinicInventory) error {
	if err := cir.db.WithContext(ctx).Create(clinicInventory).Error; err != nil {
		return err
	}
	return nil
}

func (cir *ClinicInventoryRepository) Update(ctx context.Context, clinicInventory *entity.ClinicInventory) error {
	if err := cir.db.WithContext(ctx).
		Model(clinicInventory).
		Omit("id", "clinic_id", "medicine_id").
		Updates(clinicInventory).Error; err != nil {
		return err
	}
	return nil
}

func (cir *ClinicInventoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := cir.db.WithContext(ctx).
		Delete(&entity.ClinicInventory{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
