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

func (c *clinicHandler) GetAll(ctx *gin.Context) {
	pageQuery := ctx.DefaultQuery("page", "1")
	page := utils.StringToInt(pageQuery)
	clinics, err := c.clinicUsecase.GetAll(ctx, page)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve clinics", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinics retrieved successfully", clinics)
	ctx.JSON(http.StatusOK, res)
}

func (c *clinicHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	clinic, err := c.clinicUsecase.GetByID(ctx, parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic retrieved successfully", clinic)
	ctx.JSON(http.StatusOK, res)
}

func (c *clinicHandler) Find(ctx *gin.Context) {
	name := ctx.Query("name")
	pageQuery := ctx.Query("page")
	page := utils.StringToInt(pageQuery)
	clinics, err := c.clinicUsecase.Find(ctx, name, page)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to find clinics", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinics found successfully", clinics)
	ctx.JSON(http.StatusOK, res)
}

func (c *clinicHandler) Create(ctx *gin.Context) {
	var req dto.ClinicCreateRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	clinic, err := c.clinicUsecase.Create(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic created successfully", clinic)
	ctx.JSON(http.StatusCreated, res)
}

func (c *clinicHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.ClinicUpdateRequestDTO
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

	err = c.clinicUsecase.Update(ctx, parsedID, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic updated successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *clinicHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.clinicUsecase.Delete(ctx, parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete clinic", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic deleted successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}

func (c *clinicHandler) AssignDoctor(ctx *gin.Context) {
	var req dto.AssignDoctorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.clinicUsecase.AssignDoctor(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to assign doctor", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Doctor assigned successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *clinicHandler) RemoveDoctor(ctx *gin.Context) {
	var req dto.RemoveDoctorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.clinicUsecase.RemoveDoctor(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to unassign doctor", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Doctor unassigned successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}
