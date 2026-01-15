package usecase

import (
	"context"

	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"
	"MediLink/internal/utils"

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

func (u *MedicineUsecase) GetAll(ctx context.Context, pageStr string, limitStr string) (dto.MedicineSearchResponse, error) {
	limit, err := utils.StringToInt(limitStr)
	if err != nil {
		limit = constants.PAGE_LIMIT_DEFAULT
	}

	page, err := utils.StringToInt(pageStr)
	if err != nil {
		page = 1
	}

	offset := (page - 1) * limit
	medicines, count, err := u.medicineRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dto.MedicineSearchResponse{}, err
	}

	metadata := dto.NewMetadata(int64(page), int64(limit), count)
	result := dto.ToMedicineSearchResponse(medicines, metadata)
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

func (u *MedicineUsecase) Search(ctx context.Context, name string, pageStr string, limitStr string) (dto.MedicineSearchResponse, error) {
	limit, err := utils.StringToInt(limitStr)
	if err != nil {
		limit = constants.PAGE_LIMIT_DEFAULT
	}

	page, err := utils.StringToInt(pageStr)
	if err != nil {
		page = 1
	}

	offset := (page - 1) * limit
	medicines, count, err := u.medicineRepo.Search(ctx, name, limit, offset)
	if err != nil {
		return dto.MedicineSearchResponse{}, err
	}

	metadata := dto.NewMetadata(int64(page), int64(limit), count)
	result := dto.ToMedicineSearchResponse(medicines, metadata)
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
