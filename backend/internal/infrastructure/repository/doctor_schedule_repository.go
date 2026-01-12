package repository

import (
	"context"
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type doctorScheduleRepository struct {
	db *gorm.DB
}

func NewDoctorScheduleRepository(db *gorm.DB) repository.DoctorScheduleRepository {
	return &doctorScheduleRepository{db: db}
}

func (r *doctorScheduleRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.DoctorSchedule, error) {
	var schedule entity.DoctorSchedule
	if err := r.db.WithContext(ctx).
		First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *doctorScheduleRepository) GetByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]entity.DoctorSchedule, error) {
	var schedules []entity.DoctorSchedule
	if err := r.db.WithContext(ctx).
		Order("start_time asc").
		Where("doctor_id = ?", doctorID).
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *doctorScheduleRepository) GetSchedulesByDate(ctx context.Context, doctorID uuid.UUID, date time.Time) ([]entity.DoctorSchedule, error) {
	var schedules []entity.DoctorSchedule
	if err := r.db.WithContext(ctx).
		Order("start_time asc").
		Where("doctor_id = ?", doctorID).
		Where("date = ?", date).
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *doctorScheduleRepository) GetSchedulesByDay(ctx context.Context, doctorID uuid.UUID, day string) ([]entity.DoctorSchedule, error) {
	var schedules []entity.DoctorSchedule
	if err := r.db.WithContext(ctx).
		Order("start_time asc").
		Where("doctor_id = ?", doctorID).
		Where("day_of_week = ?", day).
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *doctorScheduleRepository) Create(ctx context.Context, schedule *entity.DoctorSchedule) (*entity.DoctorSchedule, error) {
	if err := r.db.WithContext(ctx).
		Create(schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *doctorScheduleRepository) Update(ctx context.Context, schedule *entity.DoctorSchedule) error {
	if err := r.db.WithContext(ctx).
		Model(schedule).
		Omit("id", "doctor_id").
		Updates(schedule).Error; err != nil {
		return err
	}
	return nil
}

func (r *doctorScheduleRepository) UpdateStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return r.db.WithContext(ctx).
		Model(&entity.DoctorSchedule{}).
		Where("id = ?", id).
		Update("is_active", isActive).
		Error
}

func (r *doctorScheduleRepository) Delete(ctx context.Context, id uuid.UUID, doctorID uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&entity.DoctorSchedule{}, "id = ? AND doctor_id = ?", id, doctorID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
