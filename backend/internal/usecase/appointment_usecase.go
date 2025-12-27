package usecase

import (
	"context"
	"errors"

	"MediLink/internal/domain/entity"
	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/helpers/constants"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type AppointmentUsecase struct {
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentUseCase(appointmentRepo repository.AppointmentRepository) usecase.AppointmentUsecase {
	return &AppointmentUsecase{
		appointmentRepo: appointmentRepo,
	}
}

func (u *AppointmentUsecase) GetAll(ctx context.Context, page int) ([]dto.AppointmentDetailResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	appointments, err := u.appointmentRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	appointmentResponses := dto.ToListAppointmentDetailResponse(appointments)
	return appointmentResponses, nil
}

func (u *AppointmentUsecase) GetDetailByID(ctx context.Context, appointmentID uuid.UUID) (dto.AppointmentDetailResponse, error) {
	appointment, err := u.appointmentRepo.GetByID(ctx, appointmentID)
	if err != nil {
		return dto.AppointmentDetailResponse{}, err
	}

	appointmentResponse := dto.ToAppointmentDetailResponse(appointment)
	return *appointmentResponse, nil
}

func (u *AppointmentUsecase) GetByDoctor(ctx context.Context, doctorID uuid.UUID, page int) ([]dto.AppointmentDetailResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	appointments, err := u.appointmentRepo.GetByDoctorID(ctx, doctorID, limit, offset)
	if err != nil {
		return nil, err
	}

	appointmentResponses := dto.ToListAppointmentDetailResponse(appointments)
	return appointmentResponses, nil
}

func (u *AppointmentUsecase) GetByPatient(ctx context.Context, patientID uuid.UUID, page int) ([]dto.AppointmentDetailResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	appointments, err := u.appointmentRepo.GetByPatientID(ctx, patientID, limit, offset)
	if err != nil {
		return nil, err
	}

	appointmentResponses := dto.ToListAppointmentDetailResponse(appointments)
	return appointmentResponses, nil
}

func (u *AppointmentUsecase) Create(ctx context.Context, request dto.AppointmentCreateRequest) (dto.AppointmentDetailResponse, error) {
	appointment := &entity.Appointment{}
	request.ToModel(appointment)
	if err := u.appointmentRepo.Create(ctx, appointment); err != nil {
		return dto.AppointmentDetailResponse{}, err
	}

	appointmentResponse := dto.ToAppointmentDetailResponse(appointment)
	return *appointmentResponse, nil
}

func (u *AppointmentUsecase) CancelBooking(ctx context.Context, appointmentID uuid.UUID) error {
	appointment, err := u.appointmentRepo.GetByID(ctx, appointmentID)
	if err != nil {
		return err
	}

	if appointment.Status == enum.AppointmentCompleted {
		return errors.New("Appointment have already completed can not be canceled")
	}

	if err := u.appointmentRepo.UpdateStatus(ctx, appointmentID, enum.AppointmentCanceled); err != nil {
		return err
	}
	return nil
}

func (u *AppointmentUsecase) CompleteConsultation(ctx context.Context, appointmentID uuid.UUID) error {
	if err := u.appointmentRepo.UpdateStatus(ctx, appointmentID, enum.AppointmentCompleted); err != nil {
		return err
	}
	return nil
}

func (u *AppointmentUsecase) Delete(ctx context.Context, appointmentID uuid.UUID) error {
	if err := u.appointmentRepo.Delete(ctx, appointmentID); err != nil {
		return err
	}
	return nil
}
