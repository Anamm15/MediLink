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
	clinicRepo                  repository.ClinicRepository
	doctorClinicReplacementRepo repository.DoctorClinicPlacementRepository
}

func NewClinicUsecase(
	clinicRepo repository.ClinicRepository,
	doctorClinicReplacementRepo repository.DoctorClinicPlacementRepository,
) usecase.ClinicUsecase {
	return &clinicUsecase{
		clinicRepo:                  clinicRepo,
		doctorClinicReplacementRepo: doctorClinicReplacementRepo,
	}
}

func (u *clinicUsecase) GetAll(ctx context.Context, page int) ([]dto.ClinicResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	clinic, err := u.clinicRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var clinicDTOs []dto.ClinicResponse
	for _, clinic := range clinic {
		clinicDTOs = append(clinicDTOs, dto.ToClinicResponse(&clinic))
	}
	return clinicDTOs, nil
}

func (u *clinicUsecase) GetByID(ctx context.Context, id uuid.UUID) (dto.ClinicResponse, error) {
	clinic, err := u.clinicRepo.GetByID(ctx, id)
	if err != nil {
		return dto.ClinicResponse{}, err
	}
	return dto.ToClinicResponse(clinic), nil
}

func (u *clinicUsecase) Find(ctx context.Context, name string, page int) ([]dto.ClinicResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit

	clinic, err := u.clinicRepo.Find(ctx, name, limit, offset)
	if err != nil {
		return nil, err
	}
	var clinicDTOs []dto.ClinicResponse
	for _, clinic := range clinic {
		clinicDTOs = append(clinicDTOs, dto.ToClinicResponse(&clinic))
	}
	return clinicDTOs, nil
}

func (u *clinicUsecase) Create(ctx context.Context, request dto.ClinicCreateRequest) (dto.ClinicResponse, error) {
	var clinicEntity entity.Clinic
	request.ToModel(&clinicEntity)

	createdClinic, err := u.clinicRepo.Create(ctx, &clinicEntity)
	if err != nil {
		return dto.ClinicResponse{}, err
	}

	return dto.ToClinicResponse(createdClinic), nil
}

func (u *clinicUsecase) Update(ctx context.Context, id uuid.UUID, request dto.ClinicUpdateRequest) error {
	clinic, err := u.clinicRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	request.ToModel(clinic)
	return u.clinicRepo.Update(ctx, clinic)
}

func (u *clinicUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.clinicRepo.Delete(ctx, id)
}

func (u *clinicUsecase) AssignDoctor(ctx context.Context, request dto.AssignDoctorRequest) error {
	doctorClinicModel := &entity.DoctorClinicPlacement{}
	request.ToModel(doctorClinicModel)
	return u.doctorClinicReplacementRepo.Add(ctx, doctorClinicModel)
}

func (u *clinicUsecase) RemoveDoctor(ctx context.Context, request dto.RemoveDoctorRequest) error {
	return u.doctorClinicReplacementRepo.Delete(ctx, request.DoctorID, request.ClinicID)
}
