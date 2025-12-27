package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type MedicalRecordUsecase interface {
	GetByDoctor(ctx context.Context, doctorID uuid.UUID) ([]dto.MedicalRecordResponse, error)
	GetByPatient(ctx context.Context, patientID uuid.UUID) ([]dto.MedicalRecordResponse, error)
	GetById(ctx context.Context, id uuid.UUID) (dto.MedicalRecordResponse, error)
	Create(ctx context.Context, userID uuid.UUID, request *dto.MedicalRecordCreateRequest) (dto.MedicalRecordResponse, error)
	Update(ctx context.Context, id uuid.UUID, userID uuid.UUID, request *dto.MedicalRecordUpdateRequest) (dto.MedicalRecordResponse, error)
	Delete(ctx context.Context, id uuid.UUID, doctorID uuid.UUID) error
}
