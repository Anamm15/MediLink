package usecase

import (
	"context"
	"fmt"
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type patientUsecase struct {
	patientRepository repository.PatientRepository
	cacheRepository   repository.CacheRepository
}

func NewPatientUsecase(
	patientRepository repository.PatientRepository,
	cacheRepository repository.CacheRepository,
) usecase.PatientUsecase {
	return &patientUsecase{
		patientRepository: patientRepository,
		cacheRepository:   cacheRepository,
	}
}

func (u *patientUsecase) Me(ctx context.Context, userID uuid.UUID) (dto.PatientResponse, error) {
	patient, err := u.patientRepository.GetByUserID(ctx, userID)
	if err != nil {
		return dto.PatientResponse{}, err
	}
	return dto.ToPatientResponse(patient), nil
}

func (u *patientUsecase) Update(ctx context.Context, userID uuid.UUID, request dto.PatientUpdateRequest) error {
	key := fmt.Sprintf(constants.RedisKeyPatient, userID.String())
	var patient *entity.Patient

	patientIDStr, err := u.cacheRepository.Get(ctx, key)
	if err == nil {
		patientID, _ := uuid.Parse(patientIDStr)
		patient, err = u.patientRepository.GetByID(ctx, patientID)
		if err != nil {
			return err
		}
	} else {
		patient, err = u.patientRepository.GetByUserID(ctx, userID)
		if err != nil {
			return err
		}

		_ = u.cacheRepository.Set(
			ctx,
			key,
			patient.ID.String(),
			time.Hour,
		)
	}

	request.ToModel(patient)
	return u.patientRepository.Update(ctx, patient)
}
