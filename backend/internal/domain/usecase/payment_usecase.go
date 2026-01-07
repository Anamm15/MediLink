package usecase

import "MediLink/internal/dto"

type PaymentUsecase interface {
	RequestPayment(request dto.PaymentGatewayRequest) (dto.PaymentGatewayResponse, error)
}
