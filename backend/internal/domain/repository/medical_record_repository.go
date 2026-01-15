package repository

import (
	"context"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type MedicalRecordRepository interface {
	GetByPatientID(ctx context.Context, patientID uuid.UUID, limit int, offset int) ([]entity.MedicalRecord, int64, error)
	GetByDoctorID(ctx context.Context, doctorID uuid.UUID, limit int, offset int) ([]entity.MedicalRecord, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.MedicalRecord, error)
	Create(ctx context.Context, medicalRecord *entity.MedicalRecord) error
	Update(ctx context.Context, medicalRecord *entity.MedicalRecord) error
	Delete(ctx context.Context, id uuid.UUID, doctorID uuid.UUID) error
}
