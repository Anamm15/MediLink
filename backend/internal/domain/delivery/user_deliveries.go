package delivery

import "github.com/gin-gonic/gin"

type UserDelivery interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	// RefreshToken(ctx *gin.Context)
	// Logout(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	Delete(ctx *gin.Context)
	SendOTP(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
	OnBoardPatient(ctx *gin.Context)
	// ApplyAsDoctor(ctx *gin.Context)
	// ApplyAsStaff(ctx *gin.Context)
}
