package usecase

import (
	"context"

	"MediLink/internal/dto"
)

type PaymentUsecase interface {
	RequestPayment(request dto.PaymentGatewayRequest) (dto.PaymentGatewayResponse, error)
	ReceiveNotification(ctx context.Context, request dto.PaymentGatewayCallbackRequest) error
}
