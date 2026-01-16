package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type AppointmentUsecase interface {
	GetAll(ctx context.Context, page string, limit string) (dto.AppointmentResponse, error)
	GetDetailByID(ctx context.Context, appointmentID uuid.UUID) (dto.AppointmentDetailResponse, error)
	GetByDoctor(ctx context.Context, doctorID uuid.UUID, page string, limit string, statusTab string) (dto.AppointmentResponse, error)
	GetByPatient(ctx context.Context, patientID uuid.UUID, page string, limit string, statusTab string) (dto.AppointmentResponse, error)
	CreateBooking(ctx context.Context, userID uuid.UUID, request dto.CreateBookingRequest) (dto.BookingResponse, error)
	CancelBooking(ctx context.Context, appointmentID uuid.UUID) error
	CompleteConsultation(ctx context.Context, appointmentID uuid.UUID) error
	Delete(ctx context.Context, appointmentID uuid.UUID) error
}
