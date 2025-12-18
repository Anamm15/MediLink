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

func (p *patientHandler) Update(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)

	var data dto.PatientUpdateRequestDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := p.patientUsecase.Update(ctx.Request.Context(), userId, data)
	if err != nil {
		res := utils.BuildResponseSuccess("Patient updated succesfully", nil)
		ctx.JSON(http.StatusOK, res)
		return
	}
}
