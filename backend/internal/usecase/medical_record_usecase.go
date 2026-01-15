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
	"MediLink/internal/utils"

	"github.com/google/uuid"
)

type MedicalRecordUsecase struct {
	MedicalRecordRepository repository.MedicalRecordRepository
	DoctorRepository        repository.DoctorRepository
	CacheRepository         repository.CacheRepository
}

func NewMedicalRecordUsecase(
	medicalRecordRepository repository.MedicalRecordRepository,
	doctorRepository repository.DoctorRepository,
	cacheRepository repository.CacheRepository,
) usecase.MedicalRecordUsecase {
	return &MedicalRecordUsecase{
		MedicalRecordRepository: medicalRecordRepository,
		DoctorRepository:        doctorRepository,
		CacheRepository:         cacheRepository,
	}
}

func (u *MedicalRecordUsecase) GetByDoctor(ctx context.Context, doctorID uuid.UUID, pageStr string, limitStr string) (dto.MedicalRecordSearchResponse, error) {
	limit, err := utils.StringToInt(limitStr)
	if err != nil {
		limit = constants.PAGE_LIMIT_DEFAULT
	}

	page, err := utils.StringToInt(pageStr)
	if err != nil {
		page = 1
	}

	offset := (page - 1) * limit
	medicalRecords, count, err := u.MedicalRecordRepository.GetByDoctorID(ctx, doctorID, limit, offset)
	if err != nil {
		return dto.MedicalRecordSearchResponse{}, err
	}

	metadata := dto.NewMetadata(int64(page), int64(limit), count)
	results := dto.ToMedicalRecordSearchResponse(medicalRecords, metadata)
	return results, nil
}

func (u *MedicalRecordUsecase) GetByPatient(ctx context.Context, patientID uuid.UUID, pageStr string, limitStr string) (dto.MedicalRecordSearchResponse, error) {
	limit, err := utils.StringToInt(limitStr)
	if err != nil {
		limit = constants.PAGE_LIMIT_DEFAULT
	}

	page, err := utils.StringToInt(pageStr)
	if err != nil {
		page = 1
	}

	offset := (page - 1) * limit
	medicalRecords, count, err := u.MedicalRecordRepository.GetByPatientID(ctx, patientID, limit, offset)
	if err != nil {
		return dto.MedicalRecordSearchResponse{}, err
	}

	metadata := dto.NewMetadata(int64(page), int64(limit), count)
	results := dto.ToMedicalRecordSearchResponse(medicalRecords, metadata)
	return results, nil
}

func (u *MedicalRecordUsecase) GetById(ctx context.Context, id uuid.UUID) (dto.MedicalRecordResponse, error) {
	medicalRecords, err := u.MedicalRecordRepository.GetByID(ctx, id)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	return dto.ToMedicalRecordResponse(medicalRecords), nil
}

func (u *MedicalRecordUsecase) Create(ctx context.Context, userID uuid.UUID, request *dto.MedicalRecordCreateRequest) (dto.MedicalRecordResponse, error) {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	var doctorID uuid.UUID
	medicalRecordEntity := &entity.MedicalRecord{}

	doctorIDStr, err := u.CacheRepository.Get(ctx, key)
	if err == nil {
		doctorID, _ = uuid.Parse(doctorIDStr)
		medicalRecordEntity.DoctorID = doctorID
	} else {
		doctor, err := u.DoctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return dto.MedicalRecordResponse{}, err
		}
		medicalRecordEntity.DoctorID = doctor.ID
		_ = u.CacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	}

	request.ToModel(medicalRecordEntity)
	err = u.MedicalRecordRepository.Create(ctx, medicalRecordEntity)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	return dto.ToMedicalRecordResponse(medicalRecordEntity), nil
}

func (u *MedicalRecordUsecase) Update(ctx context.Context, id uuid.UUID, userID uuid.UUID, request *dto.MedicalRecordUpdateRequest) (dto.MedicalRecordResponse, error) {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	var doctorID uuid.UUID
	medicalRecordEntity := &entity.MedicalRecord{}

	doctorIDStr, err := u.CacheRepository.Get(ctx, key)
	if err == nil {
		doctorID, _ = uuid.Parse(doctorIDStr)
		medicalRecordEntity.DoctorID = doctorID
	} else {
		doctor, err := u.DoctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return dto.MedicalRecordResponse{}, err
		}
		medicalRecordEntity.DoctorID = doctor.ID
		_ = u.CacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	}

	medicalRecordEntity, err = u.MedicalRecordRepository.GetByID(ctx, id)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	if medicalRecordEntity.DoctorID != doctorID {
		return dto.MedicalRecordResponse{}, fmt.Errorf("you are not allowed to update this medical record")
	}

	request.ToModel(medicalRecordEntity)

	err = u.MedicalRecordRepository.Update(ctx, medicalRecordEntity)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	return dto.ToMedicalRecordResponse(medicalRecordEntity), nil
}

func (u *MedicalRecordUsecase) Delete(ctx context.Context, id uuid.UUID, doctorID uuid.UUID) error {
	return u.MedicalRecordRepository.Delete(ctx, id, doctorID)
}
