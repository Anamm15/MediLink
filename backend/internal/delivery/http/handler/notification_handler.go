package handler

import (
	"MediLink/internal/domain/delivery/http/handler"
	"MediLink/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notificationUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) handler.NotificationHandler {
	return &NotificationHandler{notificationUsecase: notificationUsecase}
}

func (h *NotificationHandler) SendOTPViaEmail(ctx *gin.Context) {}
