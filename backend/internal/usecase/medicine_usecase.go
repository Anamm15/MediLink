package usecase

import (
	"context"

	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type MedicineUsecase struct {
	medicineRepo repository.MedicineRepository
}

func NewMedicineUsecase(medicineRepo repository.MedicineRepository) usecase.MedicineUsecase {
	return &MedicineUsecase{
		medicineRepo: medicineRepo,
	}
}

func (mu *MedicineUsecase) GetAll(ctx context.Context, page int) ([]dto.MedicineResponseDTO, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	medicines, err := mu.medicineRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	result := dto.ToListMedicineResponseDTO(medicines)
	return result, nil
}

func (mu *MedicineUsecase) GetByID(ctx context.Context, id uuid.UUID) (dto.MedicineResponseDTO, error) {
	medicine, err := mu.medicineRepo.GetByID(ctx, id)
	if err != nil {
		return dto.MedicineResponseDTO{}, err
	}
	response := dto.ToMedicineResponseDTO(medicine)
	return *response, nil
}

func (mu *MedicineUsecase) Search(ctx context.Context, name string, page int) ([]dto.MedicineResponseDTO, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	medicines, err := mu.medicineRepo.Search(ctx, name, limit, offset)
	if err != nil {
		return nil, err
	}
	result := dto.ToListMedicineResponseDTO(medicines)
	return result, nil
}

func (mu *MedicineUsecase) Create(ctx context.Context, data dto.MedicineCreateDTO) (dto.MedicineResponseDTO, error) {
	medicine := data.ToModel()

	err := mu.medicineRepo.Create(ctx, medicine)
	if err != nil {
		return dto.MedicineResponseDTO{}, err
	}

	response := dto.ToMedicineResponseDTO(medicine)
	return *response, nil
}

func (mu *MedicineUsecase) Update(ctx context.Context, id uuid.UUID, data *dto.MedicineUpdateDTO) (dto.MedicineResponseDTO, error) {
	medicine, err := mu.medicineRepo.GetByID(ctx, id)
	if err != nil {
		return dto.MedicineResponseDTO{}, err
	}

	updatedMedicine := data.ToModel(medicine)

	err = mu.medicineRepo.Update(ctx, updatedMedicine)
	if err != nil {
		return dto.MedicineResponseDTO{}, err
	}

	response := dto.ToMedicineResponseDTO(updatedMedicine)
	return *response, nil
}

func (mu *MedicineUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	err := mu.medicineRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
