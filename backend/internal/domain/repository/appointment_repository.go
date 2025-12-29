package repository

import (
	"context"
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type AppointmentRepository interface {
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Appointment, error)
	GetByID(ctx context.Context, appointmentID uuid.UUID) (*entity.Appointment, error)
	GetByDate(ctx context.Context, date time.Time) ([]entity.Appointment, error)
	GetByDoctorID(ctx context.Context, doctorID uuid.UUID, limit int, offset int) ([]entity.Appointment, error)
	GetByPatientID(ctx context.Context, patientID uuid.UUID, limit int, offset int) ([]entity.Appointment, error)
	Create(ctx context.Context, appointment *entity.Appointment) error
	UpdateStatus(ctx context.Context, appointmentID uuid.UUID, status enum.AppointmentStatus) error
	Delete(ctx context.Context, appointmentID uuid.UUID) error
}
