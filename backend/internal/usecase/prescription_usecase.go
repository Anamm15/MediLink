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

type PrescriptionUsecase struct {
	prescriptionRepository     repository.PrescriptionRepository
	prescriptionItemRepository repository.PrescriptionItemRepository
	doctorRepository           repository.DoctorRepository
	cacheRepository            repository.CacheRepository
}

func NewPrescriptionUsecase(
	prescriptionRepository repository.PrescriptionRepository,
	prescriptionItemRepository repository.PrescriptionItemRepository,
	doctorRepository repository.DoctorRepository,
	cacheRepository repository.CacheRepository,
) usecase.PrescriptionUsecase {
	return &PrescriptionUsecase{
		prescriptionRepository:     prescriptionRepository,
		prescriptionItemRepository: prescriptionItemRepository,
		doctorRepository:           doctorRepository,
		cacheRepository:            cacheRepository,
	}
}

func (u *PrescriptionUsecase) GetByPatient(ctx context.Context, patientID uuid.UUID) ([]dto.PrescriptionResponse, error) {
	prescriptions, err := u.prescriptionRepository.GetByPatient(ctx, patientID)
	if err != nil {
		return nil, err
	}
	return dto.ToListPrescriptionResponseDTO(prescriptions), nil
}

func (u *PrescriptionUsecase) GetByDoctor(ctx context.Context, doctorID uuid.UUID) ([]dto.PrescriptionResponse, error) {
	prescriptions, err := u.prescriptionRepository.GetByDoctor(ctx, doctorID)
	if err != nil {
		return nil, err
	}
	return dto.ToListPrescriptionResponseDTO(prescriptions), nil
}

func (u *PrescriptionUsecase) GetDetailByID(ctx context.Context, id uuid.UUID) (*dto.PrescriptionResponse, error) {
	prescription, err := u.prescriptionRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToPrescriptionResponse(prescription), nil
}

func (u *PrescriptionUsecase) Create(ctx context.Context, userID uuid.UUID, request *dto.PrescriptionCreate) (dto.PrescriptionResponse, error) {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	var doctorID uuid.UUID
	prescription := &entity.Prescription{}

	doctorIDStr, err := u.cacheRepository.Get(ctx, key)
	if err == nil {
		doctorID, _ = uuid.Parse(doctorIDStr)
		prescription.DoctorID = doctorID
	} else {
		doctor, err := u.doctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return dto.PrescriptionResponse{}, err
		}
		prescription.DoctorID = doctor.ID
		_ = u.cacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	}

	request.ToModel(prescription)

	err = u.prescriptionRepository.Create(ctx, prescription)
	if err != nil {
		return dto.PrescriptionResponse{}, err
	}
	return *dto.ToPrescriptionResponse(prescription), nil
}

func (u *PrescriptionUsecase) Update(ctx context.Context, id uuid.UUID, userID uuid.UUID, request *dto.PrescriptionUpdate) (dto.PrescriptionResponse, error) {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	var doctorID uuid.UUID
	prescription := &entity.Prescription{}

	doctorIDStr, err := u.cacheRepository.Get(ctx, key)
	if err == nil {
		doctorID, _ = uuid.Parse(doctorIDStr)
		prescription.DoctorID = doctorID
	} else {
		doctor, err := u.doctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return dto.PrescriptionResponse{}, err
		}
		prescription.DoctorID = doctor.ID
		_ = u.cacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	}

	prescription, err = u.prescriptionRepository.GetByID(ctx, id)
	if err != nil {
		return dto.PrescriptionResponse{}, err
	}

	if prescription.DoctorID != doctorID {
		return dto.PrescriptionResponse{}, fmt.Errorf("you are not allowed to update this prescription")
	}

	request.ToModel(prescription)
	err = u.prescriptionRepository.Update(ctx, prescription)
	if err != nil {
		return dto.PrescriptionResponse{}, err
	}
	return *dto.ToPrescriptionResponse(prescription), nil
}

func (u *PrescriptionUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.prescriptionRepository.Delete(ctx, id)
}

func (u *PrescriptionUsecase) AddMedicine(ctx context.Context, prescriptionID uuid.UUID, request *dto.PrescriptionItemCreate) (dto.PrescriptionItemResponse, error) {
	prescription := &entity.PrescriptionItem{}
	request.ToModel(prescription)
	err := u.prescriptionItemRepository.Add(ctx, prescription)
	if err != nil {
		return dto.PrescriptionItemResponse{}, err
	}
	return dto.ToPrescriptionItemResponse(prescription), nil
}

func (u *PrescriptionUsecase) UpdateMedicine(ctx context.Context, id uuid.UUID, quantity int) error {
	return u.prescriptionItemRepository.Update(ctx, id, quantity)
}

func (u *PrescriptionUsecase) RemoveMedicine(ctx context.Context, id uuid.UUID) error {
	return u.prescriptionItemRepository.Delete(ctx, id)
}
