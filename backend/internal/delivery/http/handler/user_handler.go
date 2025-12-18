package handler

import (
	"net/http"

	"MediLink/internal/domain/delivery/http/handler"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) handler.UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (ud *userHandler) Register(ctx *gin.Context) {
	var req dto.UserRegistrationRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdUser, err := ud.userUsecase.Register(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Registration failed", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Registration successful", createdUser)
	ctx.JSON(http.StatusCreated, res)
}

func (ud *userHandler) Login(ctx *gin.Context) {
	var req dto.UserLoginRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := ud.userUsecase.Login(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Login failed", err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	res := utils.BuildResponseSuccess("Login successful", token)
	ctx.JSON(http.StatusOK, res)
}

func (ud *userHandler) GetAll(ctx *gin.Context) {
	pageQuery := ctx.DefaultQuery("page", "1")
	page := utils.StringToInt(pageQuery)
	users, err := ud.userUsecase.GetAll(ctx, page)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve users", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Users retrieved successfully", users)
	ctx.JSON(http.StatusOK, res)
}

func (ud *userHandler) GetProfile(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	profile, err := ud.userUsecase.GetProfile(ctx, userId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve profile", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Profile retrieved successfully", profile)
	ctx.JSON(http.StatusOK, res)
}

func (ud *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.UserUpdateProfileRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := ud.userUsecase.UpdateProfile(ctx, userId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update profile", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Profile updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (ud *userHandler) ChangePassword(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.UserChangePasswordRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := ud.userUsecase.ChangePassword(ctx, userId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to change password", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Password changed successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (ud *userHandler) Delete(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	err := ud.userUsecase.Delete(ctx, userId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete user", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("User deleted successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}

func (ud *userHandler) SendOTP(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	err := ud.userUsecase.SendOTP(ctx, userId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to send OTP", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("OTP sent successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (ud *userHandler) VerifyOTP(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.UserVerifyOTPRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := ud.userUsecase.VerifyOTP(ctx, userId, req.OTP)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to verify OTP", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("OTP verified successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (ud *userHandler) OnBoardPatient(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.PatientCreateRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Bad request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := ud.userUsecase.OnBoardPatient(ctx, userId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to on board patient", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Patient on boarded successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
