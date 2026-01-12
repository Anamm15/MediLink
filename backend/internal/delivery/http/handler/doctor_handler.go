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

func (h *doctorHandler) Me(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)
	profile, err := h.doctorUsecase.Me(ctx, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve profile", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Profile retrieved successfully", profile)
	ctx.JSON(http.StatusOK, res)
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

func (h *doctorHandler) GetSchedules(ctx *gin.Context) {
	doctorIDQuery := ctx.Query("doctor_id")
	doctorID, err := uuid.Parse(doctorIDQuery)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid doctor ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	schedules, err := h.doctorUsecase.GetDoctorSchedules(ctx, doctorID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve schedules", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Schedules retrieved successfully", schedules)
	ctx.JSON(http.StatusOK, res)
}

func (h *doctorHandler) GetAvailableSchedules(ctx *gin.Context) {
	doctorIDQuery := ctx.Query("doctor_id")
	date := ctx.Query("date")
	day := ctx.Query("day")
	doctorID, err := uuid.Parse(doctorIDQuery)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid doctor ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	schedules, err := h.doctorUsecase.GetAvailableSchedules(ctx, doctorID, date, day)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve available schedules", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Available schedules retrieved successfully", schedules)
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

func (h *doctorHandler) UpdateScheduleStatus(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)
	scheduleIDParam := ctx.Param("id")
	scheduleID, err := uuid.Parse(scheduleIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid schedule ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.DoctorUpdateStatusScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.doctorUsecase.UpdateStatusSchedule(ctx, userID, scheduleID, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update schedule status", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Schedule status updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *doctorHandler) DeleteSchedule(ctx *gin.Context) {
	scheduleID := ctx.Param("id")
	userID := ctx.MustGet("user_id").(uuid.UUID)
	parsedID, err := uuid.Parse(scheduleID)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid schedule ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.doctorUsecase.DeleteSchedule(ctx, userID, parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete schedule", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Schedule deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
