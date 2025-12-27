package usecase

import (
	"context"
	"errors"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type doctorUsecase struct {
	doctorRepo         repository.DoctorRepository
	doctorScheduleRepo repository.DoctorScheduleRepository
}

func NewDoctorUsecase(
	doctorRepo repository.DoctorRepository,
	doctorScheduleRepo repository.DoctorScheduleRepository,
) usecase.DoctorUsecase {
	return &doctorUsecase{
		doctorRepo:         doctorRepo,
		doctorScheduleRepo: doctorScheduleRepo,
	}
}

func (u *doctorUsecase) GetProfile(ctx context.Context, userID uuid.UUID) (dto.DoctorProfileResponse, error) {
	doctor, err := u.doctorRepo.GetWithSchedule(ctx, userID)
	if err != nil {
		return dto.DoctorProfileResponse{}, err
	}
	return dto.ToDoctorResponse(doctor), nil
}

func (u *doctorUsecase) Find(ctx context.Context, name string, page int) ([]dto.DoctorProfileResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	doctors, err := u.doctorRepo.Find(ctx, name, limit, offset)
	if err != nil {
		return nil, err
	}

	var results []dto.DoctorProfileResponse
	for _, doctor := range doctors {
		results = append(results, dto.ToDoctorResponse(&doctor))
	}
	return results, nil
}

func (u *doctorUsecase) Update(ctx context.Context, userID uuid.UUID, doctorID uuid.UUID, data dto.DoctorUpdateRequest) error {
	doctor, err := u.doctorRepo.GetByID(ctx, doctorID)
	if err != nil {
		return err
	}

	if doctor.UserID != userID {
		return errors.New("Your id is not match")
	}

	data.ToModel(doctor)
	return u.doctorRepo.Update(ctx, doctor)
}

func (u *doctorUsecase) AddSchedule(ctx context.Context, data dto.DoctorCreateScheduleRequest) (dto.DoctorScheduleResponse, error) {
	schedule := &entity.DoctorSchedule{}
	data.ToModel(schedule)
	createdSchedule, err := u.doctorScheduleRepo.Create(ctx, schedule)
	if err != nil {
		return dto.DoctorScheduleResponse{}, err
	}
	return dto.ToDoctorScheduleResponse(createdSchedule), nil
}

func (u *doctorUsecase) UpdateSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, data dto.DoctorUpdateScheduleRequest) error {
	schedule, err := u.doctorScheduleRepo.GetByID(ctx, scheduleID)
	if err != nil {
		return err
	}

	doctor, err := u.doctorRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if schedule.DoctorID != doctor.ID {
		return errors.New("Your id is not match")
	}

	data.ToModel(schedule)
	return u.doctorScheduleRepo.Update(ctx, schedule)
}
