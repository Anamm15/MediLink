package usecase

import (
	"context"

	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type patientUsecase struct {
	patientRepository repository.PatientRepository
}

func NewPatientUsecase(patientRepository repository.PatientRepository) usecase.PatientUsecase {
	return &patientUsecase{patientRepository: patientRepository}
}

func (u *patientUsecase) Update(ctx context.Context, patientID uuid.UUID, data dto.PatientUpdateRequest) error {
	patient, err := u.patientRepository.GetByUserID(ctx, patientID)
	if err != nil {
		return err
	}

	data.ToModel(patient)
	return u.patientRepository.Update(ctx, patient)
}
