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
	prescriptionRepository         repository.PrescriptionRepository
	prescriptionMedicineRepository repository.PrescriptionMedicineRepository
}

func NewPrescriptionUsecase(prescriptionRepository repository.PrescriptionRepository, prescriptionMedicineRepository repository.PrescriptionMedicineRepository) usecase.PrescriptionUsecase {
	return &PrescriptionUsecase{
		prescriptionRepository:         prescriptionRepository,
		prescriptionMedicineRepository: prescriptionMedicineRepository,
	}
}

func (pu *PrescriptionUsecase) GetByPatient(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponseDTO, error) {
	prescriptions, err := pu.prescriptionRepository.GetByPatient(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.ToPrescriptionListResponseDTO(prescriptions), nil
}

func (pu *PrescriptionUsecase) GetByDoctor(ctx context.Context, userID uuid.UUID) ([]dto.PrescriptionResponseDTO, error) {
	prescriptions, err := pu.prescriptionRepository.GetByDoctor(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.ToPrescriptionListResponseDTO(prescriptions), nil
}

func (pu *PrescriptionUsecase) GetDetailByID(ctx context.Context, id uuid.UUID) (*dto.PrescriptionResponseDTO, error) {
	prescription, err := pu.prescriptionRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToPrescriptionResponseDTO(prescription), nil
}

func (pu *PrescriptionUsecase) Create(ctx context.Context, data *dto.PrescriptionCreateDTO) (dto.PrescriptionResponseDTO, error) {
	prescription := &entity.Prescription{}
	data.ToModel(prescription)

	err := pu.prescriptionRepository.Create(ctx, prescription)
	if err != nil {
		return dto.PrescriptionResponseDTO{}, err
	}
	return *dto.ToPrescriptionResponseDTO(prescription), nil
}

func (pu *PrescriptionUsecase) Update(ctx context.Context, id uuid.UUID, data *dto.PrescriptionUpdateDTO) (dto.PrescriptionResponseDTO, error) {
	prescription, err := pu.prescriptionRepository.GetByID(ctx, id)
	if err != nil {
		return dto.PrescriptionResponseDTO{}, err
	}

	data.ToModel(prescription)
	err = pu.prescriptionRepository.Update(ctx, prescription)
	if err != nil {
		return dto.PrescriptionResponseDTO{}, err
	}
	return *dto.ToPrescriptionResponseDTO(prescription), nil
}

func (pu *PrescriptionUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return pu.prescriptionRepository.Delete(ctx, id)
}

func (pu *PrescriptionUsecase) AddMedicine(ctx context.Context, prescriptionID uuid.UUID, data *dto.PrescriptionMedicineCreateDTO) (dto.PrescriptionMedicineResponseDTO, error) {
	prescription := &entity.PrescriptionMedicine{}
	data.ToModel(prescription)
	err := pu.prescriptionMedicineRepository.Add(ctx, prescription)
	if err != nil {
		return dto.PrescriptionMedicineResponseDTO{}, err
	}
	return dto.ToPrescriptionMedicineResponseDTO(prescription), nil
}

func (pu *PrescriptionUsecase) UpdateMedicine(ctx context.Context, id uuid.UUID, quantity int) error {
	return pu.prescriptionMedicineRepository.Update(ctx, id, quantity)
}

func (pu *PrescriptionUsecase) RemoveMedicine(ctx context.Context, id uuid.UUID) error {
	return pu.prescriptionMedicineRepository.Delete(ctx, id)
}
