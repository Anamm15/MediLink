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

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) handler.AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.RegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdUser, err := h.authUsecase.Register(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Registration failed", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Registration successful", createdUser)
	ctx.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	accessToken, refreshToken, err := h.authUsecase.Login(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Login failed", err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	ctx.SetCookie(
		"refresh_token",
		refreshToken,
		60*60*24*7,
		"/",
		"",
		false,
		true,
	)

	res := utils.BuildResponseSuccess("Login successful", accessToken)
	ctx.JSON(http.StatusOK, res)
}

func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	refreshToken, _ := ctx.Cookie("refresh_token")

	newAccessToken, newRefreshToken, err := h.authUsecase.RefreshToken(ctx, refreshToken)
	if err != nil {
		res := utils.BuildResponseFailed("Refresh token failed", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.SetCookie(
		"refresh_token",
		newRefreshToken,
		60*60*24*7,
		"/",
		"",
		false,
		true,
	)

	res := utils.BuildResponseSuccess("Refresh token successful", newAccessToken)
	ctx.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	refreshToken, _ := ctx.Cookie("refresh_token")

	err := h.authUsecase.Logout(ctx, refreshToken)
	if err != nil {
		res := utils.BuildResponseFailed("Logout failed", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)
	res := utils.BuildResponseSuccess("Logout successful", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.authUsecase.ChangePassword(ctx, userId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to change password", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Password changed successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *AuthHandler) RequestResetPassword(ctx *gin.Context) {
	var req dto.RequestResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.authUsecase.RequestResetPassword(ctx, req.Email)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to request reset password", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Reset password request sent successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.authUsecase.ResetPassword(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to reset password", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Password reset successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
