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

type patientHandler struct {
	patientUsecase usecase.PatientUsecase
}

func NewPatientHandler(patientUsecase usecase.PatientUsecase) handler.PatientHandler {
	return &patientHandler{patientUsecase: patientUsecase}
}

func (h *patientHandler) Me(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	patient, err := h.patientUsecase.Me(ctx, userId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve patient", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Patient retrieved successfully", patient)
	ctx.JSON(http.StatusOK, res)
}

func (h *patientHandler) Update(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)

	var data dto.PatientUpdateRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.patientUsecase.Update(ctx.Request.Context(), userId, data)
	if err != nil {
		res := utils.BuildResponseSuccess("Patient updated succesfully", nil)
		ctx.JSON(http.StatusOK, res)
		return
	}
}
