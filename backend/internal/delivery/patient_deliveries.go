package delivery

import (
	"net/http"

	"MediLink/internal/domain/delivery"
	"MediLink/internal/domain/usecase"
	"MediLink/internal/dto"
	"MediLink/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type patientDelivery struct {
	patientUsecase usecase.PatientUsecase
}

func NewPatientDelivery(patientUsecase usecase.PatientUsecase) delivery.PatientDelivery {
	return &patientDelivery{patientUsecase: patientUsecase}
}

func (p *patientDelivery) Update(ctx *gin.Context) {
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
