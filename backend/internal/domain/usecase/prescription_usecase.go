package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type PrescriptionUsecase interface {
	GetByPatient(ctx context.Context, patientID uuid.UUID, page string, limit string, isRedeemed string) (dto.PrescriptionSearchResponse, error)
	GetByDoctor(ctx context.Context, doctorID uuid.UUID, page string, limit string) (dto.PrescriptionSearchResponse, error)
	GetDetailByID(ctx context.Context, id uuid.UUID) (*dto.PrescriptionResponse, error)
	Create(ctx context.Context, userID uuid.UUID, request *dto.PrescriptionCreate) (dto.PrescriptionResponse, error)
	Update(ctx context.Context, id uuid.UUID, userID uuid.UUID, request *dto.PrescriptionUpdate) (dto.PrescriptionResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
	AddMedicine(ctx context.Context, prescriptionID uuid.UUID, request *dto.PrescriptionItemCreate) (dto.PrescriptionItemResponse, error)
	UpdateMedicine(ctx context.Context, id uuid.UUID, quantity int) error
	RemoveMedicine(ctx context.Context, id uuid.UUID) error
}
