package handler

import (
	"net/http"

	"MediLink/internal/domain/delivery/http/handler"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/utils"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUsecase usecase.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase usecase.PaymentUsecase) handler.PaymentHandler {
	return &PaymentHandler{paymentUsecase: paymentUsecase}
}

func (h *PaymentHandler) ReceiveNotification(ctx *gin.Context) {
	var paymentRequest dto.PaymentGatewayCallbackRequest
	if err := ctx.ShouldBindJSON(&paymentRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.paymentUsecase.ReceiveNotification(ctx, paymentRequest)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to receive payment notification", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Payment notification received successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
