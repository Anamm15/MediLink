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

func (u *MedicineUsecase) GetAll(ctx context.Context, page int) ([]dto.MedicineResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	medicines, err := u.medicineRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	result := dto.ToListMedicineResponse(medicines)
	return result, nil
}

func (u *MedicineUsecase) GetByID(ctx context.Context, id uuid.UUID) (dto.MedicineResponse, error) {
	medicine, err := u.medicineRepo.GetByID(ctx, id)
	if err != nil {
		return dto.MedicineResponse{}, err
	}
	response := dto.ToMedicineResponse(medicine)
	return *response, nil
}

func (u *MedicineUsecase) Search(ctx context.Context, name string, page int) ([]dto.MedicineResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	medicines, err := u.medicineRepo.Search(ctx, name, limit, offset)
	if err != nil {
		return nil, err
	}
	result := dto.ToListMedicineResponse(medicines)
	return result, nil
}

func (u *MedicineUsecase) Create(ctx context.Context, request dto.MedicineCreate) (dto.MedicineResponse, error) {
	medicine := request.ToModel()

	err := u.medicineRepo.Create(ctx, medicine)
	if err != nil {
		return dto.MedicineResponse{}, err
	}

	response := dto.ToMedicineResponse(medicine)
	return *response, nil
}

func (u *MedicineUsecase) Update(ctx context.Context, id uuid.UUID, request *dto.MedicineUpdate) (dto.MedicineResponse, error) {
	medicine, err := u.medicineRepo.GetByID(ctx, id)
	if err != nil {
		return dto.MedicineResponse{}, err
	}

	updatedMedicine := request.ToModel(medicine)

	err = u.medicineRepo.Update(ctx, updatedMedicine)
	if err != nil {
		return dto.MedicineResponse{}, err
	}

	response := dto.ToMedicineResponse(updatedMedicine)
	return *response, nil
}

func (u *MedicineUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	err := u.medicineRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
