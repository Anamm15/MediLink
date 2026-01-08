package handler

import "github.com/gin-gonic/gin"

type PaymentHandler interface {
	ReceiveNotification(ctx *gin.Context)
}
