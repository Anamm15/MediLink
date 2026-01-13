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

type AppointmentHandler struct {
	appointmentUsecase usecase.AppointmentUsecase
}

func NewAppointmentHandler(appointmentUsecase usecase.AppointmentUsecase) handler.AppointmentHandler {
	return &AppointmentHandler{appointmentUsecase: appointmentUsecase}
}

func (h *AppointmentHandler) GetAll(ctx *gin.Context) {
	queryPage := ctx.Query("page")
	page := utils.StringToInt(queryPage)
	appointments, err := h.appointmentUsecase.GetAll(ctx, page)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve appointments", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointments retrieved successfully", appointments)
	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) GetDetailByID(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid appointment ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	appointment, err := h.appointmentUsecase.GetDetailByID(ctx, parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve appointment", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointment retrieved successfully", appointment)
	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) GetByDoctor(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)
	pageQuery := ctx.Query("page")
	page := utils.StringToInt(pageQuery)
	appointments, err := h.appointmentUsecase.GetByDoctor(ctx, userID, page)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve appointments", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointments retrieved successfully", appointments)
	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) GetByPatient(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)
	pageQuery := ctx.Query("page")
	page := utils.StringToInt(pageQuery)
	appointments, err := h.appointmentUsecase.GetByPatient(ctx, userID, page)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve appointments", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointments retrieved successfully", appointments)
	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) Create(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)
	var appointmentDTO dto.CreateBookingRequest
	if err := ctx.ShouldBindJSON(&appointmentDTO); err != nil {
		res := utils.BuildResponseFailed("Invalid request body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	appointment, err := h.appointmentUsecase.CreateBooking(ctx, userID, appointmentDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create appointment", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointment created successfully", appointment)
	ctx.JSON(http.StatusCreated, res)
}

func (h *AppointmentHandler) CancelBooking(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update appointment", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.appointmentUsecase.CancelBooking(ctx, id)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update appointment", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointment updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) CompleteConsultation(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update appointment", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.appointmentUsecase.CompleteConsultation(ctx, id)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update appointment", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointment updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete appointment", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.appointmentUsecase.Delete(ctx, id)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete appointment", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Appointment deleted successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}
