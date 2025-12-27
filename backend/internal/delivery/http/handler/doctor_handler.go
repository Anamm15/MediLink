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

type doctorHandler struct {
	doctorUsecase usecase.DoctorUsecase
}

func NewDoctorHandler(doctorUsecase usecase.DoctorUsecase) handler.DoctorHandler {
	return &doctorHandler{doctorUsecase: doctorUsecase}
}

func (h *doctorHandler) GetProfile(ctx *gin.Context) {
	doctorId := ctx.Param("id")

	parsedID, err := uuid.Parse(doctorId)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid doctor ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	profile, err := h.doctorUsecase.GetProfile(ctx, parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve profile", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Profile retrieved successfully", profile)
	ctx.JSON(http.StatusOK, res)
}

func (h *doctorHandler) Find(ctx *gin.Context) {
	name := ctx.Query("name")
	pageQuery := ctx.Query("page")
	page := utils.StringToInt(pageQuery)
	doctors, err := h.doctorUsecase.Find(ctx, name, page)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to find doctors", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Doctors found successfully", doctors)
	ctx.JSON(http.StatusOK, res)
}

func (h *doctorHandler) Update(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.DoctorUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.doctorUsecase.Update(ctx, userID, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update profile", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Profile updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *doctorHandler) AddSchedule(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.DoctorCreateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdSchedule, err := h.doctorUsecase.AddSchedule(ctx, userID, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to add schedule", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Schedule added successfully", createdSchedule)
	ctx.JSON(http.StatusOK, res)
}

func (h *doctorHandler) UpdateSchedule(ctx *gin.Context) {
	scheduleID := ctx.Param("id")
	userID := ctx.MustGet("user_id").(uuid.UUID)
	var req dto.DoctorUpdateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	parsedID, err := uuid.Parse(scheduleID)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid schedule ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.doctorUsecase.UpdateSchedule(ctx, userID, parsedID, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update schedule", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Schedule updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *doctorHandler) UpdateScheduleStatus(ctx *gin.Context)

func (h *doctorHandler) DeleteSchedule(ctx *gin.Context)
