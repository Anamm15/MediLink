package repository

import (
	"context"
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Appointment, int64, error)
	GetByID(ctx context.Context, appointmentID uuid.UUID) (*entity.Appointment, error)
	GetByDate(ctx context.Context, date time.Time) ([]entity.Appointment, error)
	GetByDoctorID(ctx context.Context, doctorID uuid.UUID, limit int, offset int, operator string, dateFilter *string) ([]entity.Appointment, int64, error)
	GetByPatientID(ctx context.Context, patientID uuid.UUID, limit int, offset int, operator string, dateFilter *string) ([]entity.Appointment, int64, error)
	Create(tx *gorm.DB, appointment *entity.Appointment) error
	UpdateStatus(ctx context.Context, tx *gorm.DB, appointmentID uuid.UUID, status enum.AppointmentStatus) error
	Delete(ctx context.Context, appointmentID uuid.UUID) error
	CheckAvailability(tx *gorm.DB, doctorID uuid.UUID, date time.Time, startTime string) (bool, error)
}
