package handler

import "github.com/gin-gonic/gin"

type MedicineHandler interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Search(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
