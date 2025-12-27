package handler

import (
	"github.com/gin-gonic/gin"
)

type NotificationHandler interface {
	SendOTPViaEmail(ctx *gin.Context)
}
