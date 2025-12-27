package handler

import "github.com/gin-gonic/gin"

type DoctorHandler interface {
	GetProfile(ctx *gin.Context)
	Find(ctx *gin.Context)
	Update(ctx *gin.Context)
	AddSchedule(ctx *gin.Context)
	UpdateSchedule(ctx *gin.Context)
	UpdateScheduleStatus(ctx *gin.Context)
	DeleteSchedule(ctx *gin.Context)
}
