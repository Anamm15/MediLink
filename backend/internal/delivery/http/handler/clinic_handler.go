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

type clinicHandler struct {
	clinicUsecase usecase.ClinicUsecase
}

func NewClinicHandler(clinicUsecase usecase.ClinicUsecase) handler.ClinicHandler {
	return &clinicHandler{clinicUsecase: clinicUsecase}
}

func (h *clinicHandler) GetAll(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	clinics, err := h.clinicUsecase.GetAll(ctx, page, limit)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve clinics", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinics retrieved successfully", clinics)
	ctx.JSON(http.StatusOK, res)
}

func (h *clinicHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("clinic_id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	clinic, err := h.clinicUsecase.GetByID(ctx, parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic retrieved successfully", clinic)
	ctx.JSON(http.StatusOK, res)
}

func (h *clinicHandler) Find(ctx *gin.Context) {
	name := ctx.Query("name")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	clinics, err := h.clinicUsecase.Find(ctx, name, page, limit)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to find clinics", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinics found successfully", clinics)
	ctx.JSON(http.StatusOK, res)
}

func (h *clinicHandler) Create(ctx *gin.Context) {
	var req dto.ClinicCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	clinic, err := h.clinicUsecase.Create(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic created successfully", clinic)
	ctx.JSON(http.StatusCreated, res)
}

func (h *clinicHandler) Update(ctx *gin.Context) {
	id := ctx.Param("clinic_id")
	var req dto.ClinicUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.clinicUsecase.Update(ctx, parsedID, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *clinicHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("clinic_id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.clinicUsecase.Delete(ctx, parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic deleted successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}

func (h *clinicHandler) AssignDoctor(ctx *gin.Context) {
	var req dto.AssignDoctorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.clinicUsecase.AssignDoctor(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to assign doctor", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Doctor assigned successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *clinicHandler) RemoveDoctor(ctx *gin.Context) {
	var req dto.RemoveDoctorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.clinicUsecase.RemoveDoctor(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to unassign doctor", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Doctor unassigned successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}
