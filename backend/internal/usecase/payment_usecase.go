package usecase

import (
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
)

type PaymentUsecase struct{}

func NewPaymentUsecase() usecase.PaymentUsecase {
	return &PaymentUsecase{}
}

func (u *PaymentUsecase) RequestPayment(request dto.PaymentGatewayRequest) (dto.PaymentGatewayResponse, error) {
	return dto.PaymentGatewayResponse{
		RedirectURL: "test",
	}, nil
}
