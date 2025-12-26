package usecase

import (
	"context"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type ClinicInventoryUsecase struct {
	ClinicInventoryRepository repository.ClinicInventoryRepository
}

func NewClinicInventoryUsecase(clinicInventoryRepository repository.ClinicInventoryRepository) usecase.ClinicInventoryUsecase {
	return &ClinicInventoryUsecase{
		ClinicInventoryRepository: clinicInventoryRepository,
	}
}

func (ciu *ClinicInventoryUsecase) GetByClinic(ctx context.Context, clinicID uuid.UUID) ([]dto.ClinicInventoryResponse, error) {
	inventories, err := ciu.ClinicInventoryRepository.GetByClinicID(ctx, clinicID)
	if err != nil {
		return nil, err
	}
	return dto.ToListClinicInventoryResponse(inventories), nil
}

func (ciu *ClinicInventoryUsecase) GetByID(ctx context.Context, id uuid.UUID) (dto.ClinicInventoryResponse, error) {
	inventory, err := ciu.ClinicInventoryRepository.GetByID(ctx, id)
	if err != nil {
		return dto.ClinicInventoryResponse{}, err
	}
	return *dto.ToClinicInventoryResponse(&inventory), nil
}

func (ciu *ClinicInventoryUsecase) Create(ctx context.Context, request dto.ClinicInventoryCreateRequest) (dto.ClinicInventoryResponse, error) {
	inventory := &entity.ClinicInventory{}
	request.ToModel(inventory)
	if err := ciu.ClinicInventoryRepository.Create(ctx, inventory); err != nil {
		return dto.ClinicInventoryResponse{}, err
	}
	return *dto.ToClinicInventoryResponse(inventory), nil
}

func (ciu *ClinicInventoryUsecase) Update(ctx context.Context, id uuid.UUID, request dto.ClinicInventoryUpdateRequest) (dto.ClinicInventoryResponse, error) {
	inventory, err := ciu.ClinicInventoryRepository.GetByID(ctx, id)
	if err != nil {
		return dto.ClinicInventoryResponse{}, err
	}

	request.ToModel(&inventory)
	if err := ciu.ClinicInventoryRepository.Update(ctx, &inventory); err != nil {
		return dto.ClinicInventoryResponse{}, err
	}
	return *dto.ToClinicInventoryResponse(&inventory), nil
}

func (ciu *ClinicInventoryUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	if err := ciu.ClinicInventoryRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
