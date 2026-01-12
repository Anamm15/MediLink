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
	"MediLink/internal/utils"

	"github.com/google/uuid"
)

type doctorUsecase struct {
	doctorRepository         repository.DoctorRepository
	doctorScheduleRepository repository.DoctorScheduleRepository
	appointmentRepository    repository.AppointmentRepository
	cacheRepository          repository.CacheRepository
}

func NewDoctorUsecase(
	doctorRepository repository.DoctorRepository,
	doctorScheduleRepository repository.DoctorScheduleRepository,
	appointmentRepository repository.AppointmentRepository,
	cacheRepository repository.CacheRepository,
) usecase.DoctorUsecase {
	return &doctorUsecase{
		doctorRepository:         doctorRepository,
		doctorScheduleRepository: doctorScheduleRepository,
		appointmentRepository:    appointmentRepository,
		cacheRepository:          cacheRepository,
	}
}

func (u *doctorUsecase) Me(ctx context.Context, userID uuid.UUID) (dto.DoctorProfileResponse, error) {
	doctor, err := u.doctorRepository.GetProfileByUserID(ctx, userID)
	if err != nil {
		return dto.DoctorProfileResponse{}, err
	}
	return dto.ToDoctorResponse(doctor), nil
}

func (u *doctorUsecase) GetProfile(ctx context.Context, doctorID uuid.UUID) (dto.DoctorProfileResponse, error) {
	doctor, err := u.doctorRepository.GetProfileByID(ctx, doctorID)
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

func (u *doctorUsecase) GetDoctorSchedules(ctx context.Context, doctorID uuid.UUID) ([]dto.DoctorScheduleResponse, error) {
	doctorSchedules, err := u.doctorScheduleRepository.GetByDoctorID(ctx, doctorID)
	if err != nil {
		return nil, err
	}
	return dto.ToListDoctorScheduleResponse(doctorSchedules), nil
}

func (u *doctorUsecase) GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (dto.DoctorScheduleResponse, error) {
	doctorSchedule, err := u.doctorScheduleRepository.GetByID(ctx, scheduleID)
	if err != nil {
		return dto.DoctorScheduleResponse{}, err
	}
	return dto.ToDoctorScheduleResponse(doctorSchedule), nil
}

func (u *doctorUsecase) GetAvailableSchedules(ctx context.Context, doctorID uuid.UUID, date string, day string) ([]dto.DoctorScheduleResponse, error) {
	dateTime := utils.ParseDate(date)
	reservedSchedules, err := u.appointmentRepository.GetByDate(ctx, dateTime)
	if err != nil {
		return nil, err
	}

	schedules, err := u.doctorScheduleRepository.GetSchedulesByDay(ctx, doctorID, day)
	if err != nil {
		return nil, err
	}

	var availableSchedules []dto.DoctorScheduleResponse
	for _, schedule := range schedules {
		isAvailable := true
		for _, reservedSchedule := range reservedSchedules {
			if reservedSchedule.StartTime == schedule.StartTime {
				isAvailable = false
				break
			}
		}
		if isAvailable {
			availableSchedules = append(availableSchedules, dto.ToDoctorScheduleResponse(&schedule))
		}
	}
	return availableSchedules, nil
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

func (u *doctorUsecase) UpdateStatusSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID, request dto.DoctorUpdateStatusScheduleRequest) error {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	doctorID, err := u.cacheRepository.Get(ctx, key)
	if err != nil {
		doctor, err := u.doctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return err
		}
		doctorID = doctor.ID.String()
		_ = u.cacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	}

	schedule, err := u.doctorScheduleRepository.GetByID(ctx, scheduleID)
	if err != nil {
		return err
	}

	if schedule.DoctorID.String() != doctorID {
		return errors.New("You do not have permission to update this schedule")
	}

	return u.doctorScheduleRepository.UpdateStatus(ctx, scheduleID, request.IsActive)
}

func (u *doctorUsecase) DeleteSchedule(ctx context.Context, userID uuid.UUID, scheduleID uuid.UUID) error {
	key := fmt.Sprintf(constants.RedisKeyDoctor, userID.String())
	var doctorID uuid.UUID
	doctorIDStr, err := u.cacheRepository.Get(ctx, key)
	if err != nil {
		doctor, err := u.doctorRepository.GetByUserID(ctx, userID)
		if err != nil {
			return err
		}
		doctorID = doctor.ID
		_ = u.cacheRepository.Set(
			ctx,
			key,
			doctor.ID.String(),
			time.Hour,
		)
	} else {
		doctorID, _ = uuid.Parse(doctorIDStr)
	}

	return u.doctorScheduleRepository.Delete(ctx, scheduleID, doctorID)
}
