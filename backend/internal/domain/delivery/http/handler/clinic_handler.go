package handler

import "github.com/gin-gonic/gin"

type ClinicHandler interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Find(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	AssignDoctor(ctx *gin.Context)
	RemoveDoctor(ctx *gin.Context)
}
