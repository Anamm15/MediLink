package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type PatientUsecase interface {
	Me(ctx context.Context, userID uuid.UUID) (dto.PatientResponse, error)
	Update(ctx context.Context, userID uuid.UUID, request dto.PatientUpdateRequest) error
}
