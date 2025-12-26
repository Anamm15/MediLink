package usecase

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type MedicalRecordUsecase struct {
	MedicalRecordRepository repository.MedicalRecordRepository
}

func NewMedicalRecordUsecase(medicalRecordRepository repository.MedicalRecordRepository) usecase.MedicalRecordUsecase {
	return &MedicalRecordUsecase{
		MedicalRecordRepository: medicalRecordRepository,
	}
}

func (u *MedicalRecordUsecase) GetByDoctor(ctx context.Context, doctorID uuid.UUID) ([]dto.MedicalRecordResponse, error) {
	medicalRecords, err := u.MedicalRecordRepository.GetByDoctorID(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	return dto.ToListMedicalRecordResponse(medicalRecords), nil
}

func (u *MedicalRecordUsecase) GetByPatient(ctx context.Context, patientID uuid.UUID) ([]dto.MedicalRecordResponse, error) {
	medicalRecords, err := u.MedicalRecordRepository.GetByPatientID(ctx, patientID)
	if err != nil {
		return nil, err
	}

	return dto.ToListMedicalRecordResponse(medicalRecords), nil
}

func (u *MedicalRecordUsecase) GetById(ctx context.Context, id uuid.UUID) (dto.MedicalRecordResponse, error) {
	medicalRecords, err := u.MedicalRecordRepository.GetByID(ctx, id)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	return dto.ToMedicalRecordResponse(medicalRecords), nil
}

func (u *MedicalRecordUsecase) Create(ctx context.Context, req *dto.MedicalRecordCreateRequest) (dto.MedicalRecordResponse, error) {
	medicalRecordEntity := &entity.MedicalRecord{}
	req.ToModel(medicalRecordEntity)
	err := u.MedicalRecordRepository.Create(ctx, medicalRecordEntity)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	return dto.ToMedicalRecordResponse(medicalRecordEntity), nil
}

func (u *MedicalRecordUsecase) Update(ctx context.Context, id uuid.UUID, req *dto.MedicalRecordUpdateRequest) (dto.MedicalRecordResponse, error) {
	medicalRecordEntity, err := u.MedicalRecordRepository.GetByID(ctx, id)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	req.ToModel(medicalRecordEntity)

	err = u.MedicalRecordRepository.Update(ctx, medicalRecordEntity)
	if err != nil {
		return dto.MedicalRecordResponse{}, err
	}

	return dto.ToMedicalRecordResponse(medicalRecordEntity), nil
}

func (u *MedicalRecordUsecase) Delete(ctx context.Context, id uuid.UUID, doctorID uuid.UUID) error {
	return u.MedicalRecordRepository.Delete(ctx, id, doctorID)
}
