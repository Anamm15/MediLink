package handler

import (
	"github.com/gin-gonic/gin"
)

type AppointmentHandler interface {
	GetAll(ctx *gin.Context)
	GetDetailByID(ctx *gin.Context)
	GetByDoctor(ctx *gin.Context)
	GetByPatient(ctx *gin.Context)
	Create(ctx *gin.Context)
	CancelBooking(ctx *gin.Context)
	CompleteConsultation(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
