package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type PatientUsecase interface {
	Update(ctx context.Context, userID uuid.UUID, request dto.PatientUpdateRequest) error
}
