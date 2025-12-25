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

func (au *AppointmentUsecase) GetAll(ctx context.Context, page int) ([]dto.AppointmentDetailResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	appointments, err := au.appointmentRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	appointmentResponses := dto.ToListAppointmentDetailResponse(appointments)
	return appointmentResponses, nil
}

func (au *AppointmentUsecase) GetDetailByID(ctx context.Context, appointmentID uuid.UUID) (dto.AppointmentDetailResponse, error) {
	appointment, err := au.appointmentRepo.GetByID(ctx, appointmentID)
	if err != nil {
		return dto.AppointmentDetailResponse{}, err
	}

	appointmentResponse := dto.ToAppointmentDetailResponse(appointment)
	return *appointmentResponse, nil
}

func (au *AppointmentUsecase) GetByDoctor(ctx context.Context, doctorID uuid.UUID, page int) ([]dto.AppointmentDetailResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	appointments, err := au.appointmentRepo.GetByDoctorID(ctx, doctorID, limit, offset)
	if err != nil {
		return nil, err
	}

	appointmentResponses := dto.ToListAppointmentDetailResponse(appointments)
	return appointmentResponses, nil
}

func (au *AppointmentUsecase) GetByPatient(ctx context.Context, patientID uuid.UUID, page int) ([]dto.AppointmentDetailResponse, error) {
	limit := constants.PAGE_LIMIT_DEFAULT
	offset := (page - 1) * limit
	appointments, err := au.appointmentRepo.GetByPatientID(ctx, patientID, limit, offset)
	if err != nil {
		return nil, err
	}

	appointmentResponses := dto.ToListAppointmentDetailResponse(appointments)
	return appointmentResponses, nil
}

func (au *AppointmentUsecase) Create(ctx context.Context, request dto.AppointmentCreateRequest) (dto.AppointmentDetailResponse, error) {
	appointment := &entity.Appointment{}
	request.ToModel(appointment)
	if err := au.appointmentRepo.Create(ctx, appointment); err != nil {
		return dto.AppointmentDetailResponse{}, err
	}

	appointmentResponse := dto.ToAppointmentDetailResponse(appointment)
	return *appointmentResponse, nil
}

func (au *AppointmentUsecase) CancelBooking(ctx context.Context, appointmentID uuid.UUID) error {
	appointment, err := au.appointmentRepo.GetByID(ctx, appointmentID)
	if err != nil {
		return err
	}

	if appointment.Status == enum.AppointmentCompleted {
		return errors.New("Appointment have already completed can not be canceled")
	}

	if err := au.appointmentRepo.UpdateStatus(ctx, appointmentID, enum.AppointmentCanceled); err != nil {
		return err
	}
	return nil
}

func (au *AppointmentUsecase) CompleteConsultation(ctx context.Context, appointmentID uuid.UUID) error {
	if err := au.appointmentRepo.UpdateStatus(ctx, appointmentID, enum.AppointmentCompleted); err != nil {
		return err
	}
	return nil
}

func (au *AppointmentUsecase) Delete(ctx context.Context, appointmentID uuid.UUID) error {
	if err := au.appointmentRepo.Delete(ctx, appointmentID); err != nil {
		return err
	}
	return nil
}
