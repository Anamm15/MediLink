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
	"MediLink/internal/helpers/enum"
	"MediLink/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentUsecase struct {
	db                 *gorm.DB
	appointmentRepo    repository.AppointmentRepository
	patientRepo        repository.PatientRepository
	billingRepo        repository.BillingRepository
	paymentRepo        repository.PaymentRepository
	doctorScheduleRepo repository.DoctorScheduleRepository
	cacheRepo          repository.CacheRepository
	paymentUsecase     usecase.PaymentUsecase
}

func NewAppointmentUseCase(
	db *gorm.DB,
	appointmentRepo repository.AppointmentRepository,
	patientRepo repository.PatientRepository,
	billingRepo repository.BillingRepository,
	paymentRepo repository.PaymentRepository,
	doctorScheduleRepo repository.DoctorScheduleRepository,
	cacheRepo repository.CacheRepository,
	paymentUsecase usecase.PaymentUsecase,
) usecase.AppointmentUsecase {
	return &AppointmentUsecase{
		db:                 db,
		appointmentRepo:    appointmentRepo,
		patientRepo:        patientRepo,
		billingRepo:        billingRepo,
		paymentRepo:        paymentRepo,
		doctorScheduleRepo: doctorScheduleRepo,
		cacheRepo:          cacheRepo,
		paymentUsecase:     paymentUsecase,
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

func (u *AppointmentUsecase) CreateBooking(ctx context.Context, userID uuid.UUID, req dto.CreateBookingRequest) (dto.BookingResponse, error) {
	var appointmentID uuid.UUID
	var billingID uuid.UUID
	var finalPrice float64
	var patientID uuid.UUID

	parsedDate := utils.ParseDate(req.AppointmentDate)

	// =================================================================
	// PHASE 0: RESOLVE PATIENT ID (Security & Caching)
	// =================================================================

	key := fmt.Sprintf(constants.RedisKeyPatient, userID.String())

	patientIDStr, err := u.cacheRepo.Get(ctx, key)
	if err == nil && patientIDStr != "" {
		patientID, _ = uuid.Parse(patientIDStr)
	} else {
		patient, err := u.patientRepo.GetByUserID(ctx, userID)
		if err != nil {
			return dto.BookingResponse{}, errors.New("user profile not found or not registered as patient")
		}

		patientID = patient.ID
		_ = u.cacheRepo.Set(
			ctx,
			key,
			patient.ID.String(),
			1*time.Hour,
		)
	}

	// =================================================================
	// PHASE 1: DATABASE TRANSACTION
	// =================================================================

	err = u.db.Transaction(func(tx *gorm.DB) error {
		// 1. Cek Ketersediaan Slot
		isAvailable, err := u.appointmentRepo.CheckAvailability(tx, req.DoctorID, parsedDate, req.StartTime)
		if err != nil {
			return err
		}
		if !isAvailable {
			return errors.New("Schedule is not available")
		}

		// 2. Ambil Harga Snapshot
		placement, err := u.doctorScheduleRepo.GetByID(ctx, req.ScheduleID)
		if err != nil {
			return err
		}
		finalPrice = placement.ConsultationFee

		// 3. Buat Entity Appointment
		newAppt := entity.Appointment{
			PatientID:               patientID,
			DoctorID:                req.DoctorID,
			ClinicID:                req.ClinicID,
			AppointmentDate:         parsedDate,
			StartTime:               req.StartTime,
			EndTime:                 req.EndTime,
			Status:                  enum.AppointmentPending,
			Type:                    req.Type,
			ConsultationFeeSnapshot: finalPrice,
		}

		if err := u.appointmentRepo.Create(tx, &newAppt); err != nil {
			return err
		}
		appointmentID = newAppt.ID

		// 4. Buat Entity Billing
		newBilling := entity.Billing{
			AppointmentID: &appointmentID,
			PatientID:     patientID,
			TotalAmount:   finalPrice,
		}

		if err := u.billingRepo.Create(tx, &newBilling); err != nil {
			return err
		}
		billingID = newBilling.ID

		return nil
	})
	if err != nil {
		return dto.BookingResponse{}, err
	}

	// =================================================================
	// PHASE 2: EXTERNAL PAYMENT GATEWAY
	// =================================================================

	paymentReq := dto.PaymentGatewayRequest{
		OrderID: billingID.String(),
		Amount:  finalPrice,
	}

	paymentResponse, err := u.paymentUsecase.RequestPayment(paymentReq)

	paymentUrl := ""
	if err == nil {
		paymentUrl = paymentResponse.RedirectURL
	}

	// =================================================================
	// PHASE 3: UPDATE PAYMENT DATA
	// =================================================================

	if paymentUrl != "" {
		newPayment := entity.Payment{
			BillingID:     billingID,
			Amount:        finalPrice,
			Status:        "unpaid",
			PaymentMethod: "pending_selection",
			PaymentURL:    &paymentUrl,
		}
		u.paymentRepo.Create(ctx, &newPayment)
	}

	return dto.BookingResponse{
		AppointmentID: appointmentID,
		PaymentURL:    &paymentUrl,
	}, nil
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
