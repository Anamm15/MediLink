package usecase

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type PrescriptionUsecase struct {
	prescriptionRepository     repository.PrescriptionRepository
	prescriptionItemRepository repository.PrescriptionItemRepository
}

func NewPrescriptionUsecase(
	prescriptionRepository repository.PrescriptionRepository,
	prescriptionItemRepository repository.PrescriptionItemRepository,
) usecase.PrescriptionUsecase {
	return &PrescriptionUsecase{
		prescriptionRepository:     prescriptionRepository,
		prescriptionItemRepository: prescriptionItemRepository,
	}
}

func (u *PrescriptionUsecase) GetByPatient(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponse, error) {
	prescriptions, err := u.prescriptionRepository.GetByPatient(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.ToListPrescriptionResponseDTO(prescriptions), nil
}

func (u *PrescriptionUsecase) GetByDoctor(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponse, error) {
	prescriptions, err := u.prescriptionRepository.GetByDoctor(ctx, userID)
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

func (u *PrescriptionUsecase) Create(ctx context.Context, data *dto.PrescriptionCreate) (dto.PrescriptionResponse, error) {
	prescription := &entity.Prescription{}
	data.ToModel(prescription)

	err := u.prescriptionRepository.Create(ctx, prescription)
	if err != nil {
		return dto.PrescriptionResponse{}, err
	}
	return *dto.ToPrescriptionResponse(prescription), nil
}

func (u *PrescriptionUsecase) Update(ctx context.Context, id uuid.UUID, data *dto.PrescriptionUpdate) (dto.PrescriptionResponse, error) {
	prescription, err := u.prescriptionRepository.GetByID(ctx, id)
	if err != nil {
		return dto.PrescriptionResponse{}, err
	}

	data.ToModel(prescription)
	err = u.prescriptionRepository.Update(ctx, prescription)
	if err != nil {
		return dto.PrescriptionResponse{}, err
	}
	return *dto.ToPrescriptionResponse(prescription), nil
}

func (u *PrescriptionUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.prescriptionRepository.Delete(ctx, id)
}

func (u *PrescriptionUsecase) AddMedicine(ctx context.Context, prescriptionID uuid.UUID, data *dto.PrescriptionItemCreate) (dto.PrescriptionItemResponse, error) {
	prescription := &entity.PrescriptionItem{}
	data.ToModel(prescription)
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
