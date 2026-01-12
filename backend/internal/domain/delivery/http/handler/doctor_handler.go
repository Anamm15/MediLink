package handler

import "github.com/gin-gonic/gin"

type DoctorHandler interface {
	Me(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	Find(ctx *gin.Context)
	Update(ctx *gin.Context)
	AddSchedule(ctx *gin.Context)
	GetSchedules(ctx *gin.Context)
	GetScheduleByID(ctx *gin.Context)
	GetAvailableSchedules(ctx *gin.Context)
	UpdateSchedule(ctx *gin.Context)
	UpdateScheduleStatus(ctx *gin.Context)
	DeleteSchedule(ctx *gin.Context)
}
