package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetAll(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	Me(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	Delete(ctx *gin.Context)
	SendVerificationUser(ctx *gin.Context)
	VerifyUser(ctx *gin.Context)
	OnBoardPatient(ctx *gin.Context)
}
