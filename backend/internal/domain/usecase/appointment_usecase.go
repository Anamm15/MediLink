package usecase

import (
	"context"

	"MediLink/internal/dto"

	"github.com/google/uuid"
)

type AppointmentUsecase interface {
	GetAll(ctx context.Context, page int) ([]dto.AppointmentDetailResponse, error)
	GetDetailByID(ctx context.Context, appointmentID uuid.UUID) (dto.AppointmentDetailResponse, error)
	GetByDoctor(ctx context.Context, doctorID uuid.UUID, page int) ([]dto.AppointmentDetailResponse, error)
	GetByPatient(ctx context.Context, patientID uuid.UUID, page int) ([]dto.AppointmentDetailResponse, error)
	CreateBooking(ctx context.Context, userID uuid.UUID, request dto.CreateBookingRequest) (dto.BookingResponse, error)
	CancelBooking(ctx context.Context, appointmentID uuid.UUID) error
	CompleteConsultation(ctx context.Context, appointmentID uuid.UUID) error
	Delete(ctx context.Context, appointmentID uuid.UUID) error
}
