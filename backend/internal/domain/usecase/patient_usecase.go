package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type PatientUsecase interface {
	Update(ctx context.Context, patientID uuid.UUID, data dto.PatientUpdateRequestDTO) error
}
