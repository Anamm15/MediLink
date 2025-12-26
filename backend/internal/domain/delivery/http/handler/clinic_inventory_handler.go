package handler

import "github.com/gin-gonic/gin"

type ClinicInventoryHandler interface {
	GetByClinic(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
