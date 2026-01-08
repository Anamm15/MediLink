package usecase

import (
	"context"
	"fmt"

	"MediLink/internal/domain/repository"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/utils"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type PaymentUsecase struct {
	db                    *gorm.DB
	snapClient            snap.Client
	paymentRepository     repository.PaymentRepository
	billingRepository     repository.BillingRepository
	appointmentRepository repository.AppointmentRepository
}

func NewPaymentUsecase(
	db *gorm.DB,
	serverKey string,
	paymentRepository repository.PaymentRepository,
	billingRepository repository.BillingRepository,
	appointmentRepository repository.AppointmentRepository,
) usecase.PaymentUsecase {
	var client snap.Client
	client.New(serverKey, midtrans.Sandbox)
	return &PaymentUsecase{
		db:                    db,
		snapClient:            client,
		paymentRepository:     paymentRepository,
		billingRepository:     billingRepository,
		appointmentRepository: appointmentRepository,
	}
}

func (u *PaymentUsecase) RequestPayment(request dto.PaymentGatewayRequest) (dto.PaymentGatewayResponse, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.OrderID,
			GrossAmt: request.GrossAmount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: request.Name,
			Email: request.Email,
			Phone: request.PhoneNumber,
		},
	}

	snapResp, err := u.snapClient.CreateTransaction(req)
	if err != nil {
		return dto.PaymentGatewayResponse{}, err
	}

	paymentGatewayResponse := dto.PaymentGatewayResponse{
		RedirectURL: snapResp.RedirectURL,
		Token:       snapResp.Token,
	}

	return paymentGatewayResponse, nil
}

func (u *PaymentUsecase) ReceiveNotification(ctx context.Context, request dto.PaymentGatewayCallbackRequest) error {
	orderID, err := uuid.Parse(request.OrderID)
	if err != nil {
		return fmt.Errorf("invalid order id: %w", err)
	}

	tx := u.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	billing, err := u.billingRepository.GetByID(tx, orderID)
	if err != nil {
		return fmt.Errorf("billing not found: %w", err)
	}

	paymentStatus, appointmentStatus := utils.MapMidtransStatus(request.TransactionStatus, request.FraudStatus)
	if err = u.paymentRepository.UpdateStatus(tx, orderID, paymentStatus); err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	if appointmentStatus != "" && billing.AppointmentID != nil {
		if err = u.appointmentRepository.UpdateStatus(ctx, tx, *billing.AppointmentID, appointmentStatus); err != nil {
			return fmt.Errorf("failed to update appointment status: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
