package usecase

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type clinicUsecase struct {
	clinicRepo repository.ClinicRepository
}

func NewClinicUsecase(clinicRepo repository.ClinicRepository) usecase.ClinicUsecase {
	return &clinicUsecase{
		clinicRepo: clinicRepo,
	}
}

func (c *clinicUsecase) GetAll(ctx context.Context, page int) ([]dto.ClinicResponseDTO, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	clinic, err := c.clinicRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var clinicDTOs []dto.ClinicResponseDTO
	for _, clinic := range clinic {
		clinicDTOs = append(clinicDTOs, dto.ToClinicResponseDTO(&clinic))
	}
	return clinicDTOs, nil
}

func (c *clinicUsecase) GetByID(ctx context.Context, id uuid.UUID) (dto.ClinicResponseDTO, error) {
	clinic, err := c.clinicRepo.GetByID(ctx, id)
	if err != nil {
		return dto.ClinicResponseDTO{}, err
	}
	return dto.ToClinicResponseDTO(clinic), nil
}

func (c *clinicUsecase) Find(ctx context.Context, name string, page int) ([]dto.ClinicResponseDTO, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	clinic, err := c.clinicRepo.Find(ctx, name, limit, offset)
	if err != nil {
		return nil, err
	}
	var clinicDTOs []dto.ClinicResponseDTO
	for _, clinic := range clinic {
		clinicDTOs = append(clinicDTOs, dto.ToClinicResponseDTO(&clinic))
	}
	return clinicDTOs, nil
}

func (c *clinicUsecase) Create(ctx context.Context, data dto.ClinicCreateRequestDTO) (dto.ClinicResponseDTO, error) {
	var clinicEntity entity.Clinic
	data.AssignToEntity(&clinicEntity)

	createdClinic, err := c.clinicRepo.Create(ctx, &clinicEntity)
	if err != nil {
		return dto.ClinicResponseDTO{}, err
	}

	return dto.ToClinicResponseDTO(createdClinic), nil
}

func (c *clinicUsecase) Update(ctx context.Context, id uuid.UUID, data dto.ClinicUpdateRequestDTO) error {
	clinic, err := c.clinicRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	data.AssignToEntity(clinic)
	return c.clinicRepo.Update(ctx, clinic)
}

func (c *clinicUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return c.clinicRepo.Delete(ctx, id)
}
