package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"

	"github.com/google/uuid"
)

type doctorUsecase struct {
	doctorRepository         repository.DoctorRepository
	doctorScheduleRepository repository.DoctorScheduleRepository
	cacheRepository          repository.CacheRepository
}

func NewDoctorUsecase(
	doctorRepository repository.DoctorRepository,
	doctorScheduleRepository repository.DoctorScheduleRepository,
	cacheRepository repository.CacheRepository,
) usecase.DoctorUsecase {
	return &doctorUsecase{
		doctorRepository:         doctorRepository,
		doctorScheduleRepository: doctorScheduleRepository,
		cacheRepository:          cacheRepository,
	}
}

func (u *doctorUsecase) GetProfile(ctx context.Context, userID uuid.UUID) (dto.DoctorProfileResponse, error) {
	doctor, err := u.doctorRepository.GetWithSchedule(ctx, userID)
	if err != nil {
		return dto.DoctorProfileResponse{}, err
	}
	return dto.ToDoctorResponse(doctor), nil
}

func (u *doctorUsecase) Find(ctx context.Context, name string, page int) ([]dto.DoctorProfileResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	doctors, err := u.doctorRepository.Find(ctx, name, limit, offset)
	if err != nil {
		return nil, err
	}

	var results []dto.DoctorProfileResponse
	for _, doctor := range doctors {
		results = append(results, dto.ToDoctorResponse(&doctor))
	}
	return results, nil
}

func (u *doctorUsecase) Update(ctx context.Context, userID uuid.UUID, request dto.DoctorUpdateRequest) error {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	var doctor *entity.Doctor

	doctorIDStr, err := u.cacheRepository.Get(ctx, key)
	if err == nil {
		doctorID, _ := uuid.Parse(doctorIDStr)
		doctor, err = u.doctorRepository.GetByID(ctx, doctorID)
		if err != nil {
			return err
		}
	} else {
		doctor, err := u.doctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return err
		}

		_ = u.cacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	}

	request.ToModel(doctor)
	return u.doctorRepository.Update(ctx, doctor)
}

func (u *doctorUsecase) AddSchedule(ctx context.Context, userID uuid.UUID, request dto.DoctorCreateScheduleRequest) (dto.DoctorScheduleResponse, error) {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	schedule := &entity.DoctorSchedule{}

	doctorIDStr, err := u.cacheRepository.Get(ctx, key)
	if err == nil {
		doctorID, _ := uuid.Parse(doctorIDStr)
		schedule.DoctorID = doctorID
	} else {
		doctor, err := u.doctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return dto.DoctorScheduleResponse{}, err
		}
		schedule.DoctorID = doctor.ID

		_ = u.cacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	}

	request.ToModel(schedule)
	createdSchedule, err := u.doctorScheduleRepository.Create(ctx, schedule)
	if err != nil {
		return dto.DoctorScheduleResponse{}, err
	}
	return dto.ToDoctorScheduleResponse(createdSchedule), nil
}

func (u *doctorUsecase) UpdateSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, request dto.DoctorUpdateScheduleRequest) error {
	schedule, err := u.doctorScheduleRepository.GetByID(ctx, scheduleID)
	if err != nil {
		return err
	}

	doctor, err := u.doctorRepository.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if schedule.DoctorID != doctor.ID {
		return errors.New("You do not have permission to update this schedule")
	}

	request.ToModel(schedule)
	return u.doctorScheduleRepository.Update(ctx, schedule)
}
