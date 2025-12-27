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

type ClinicInventoryHandler struct {
	clinicInventoryUsecase usecase.ClinicInventoryUsecase
}

func NewClinicInventoryHandler(clinicInventoryUsecase usecase.ClinicInventoryUsecase) handler.ClinicInventoryHandler {
	return &ClinicInventoryHandler{clinicInventoryUsecase: clinicInventoryUsecase}
}

func (h *ClinicInventoryHandler) GetByClinic(ctx *gin.Context) {
	clinicIDParam := ctx.Param("clinic_id")
	clinicID, err := uuid.Parse(clinicIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	inventories, err := h.clinicInventoryUsecase.GetByClinic(ctx, clinicID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve clinic inventories", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic inventories retrieved successfully", inventories)
	ctx.JSON(http.StatusOK, res)
}

func (h *ClinicInventoryHandler) GetByID(ctx *gin.Context) {
	inventoryIDParam := ctx.Param("id")
	inventoryID, err := uuid.Parse(inventoryIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	inventory, err := h.clinicInventoryUsecase.GetByID(ctx, inventoryID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve clinic inventory", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic inventory retrieved successfully", inventory)
	ctx.JSON(http.StatusOK, res)
}

func (h *ClinicInventoryHandler) Create(ctx *gin.Context) {
	var req dto.ClinicInventoryCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdinventory, err := h.clinicInventoryUsecase.Create(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create clinic inventory", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic inventory created successfully", createdinventory)
	ctx.JSON(http.StatusCreated, res)
}

func (h *ClinicInventoryHandler) Update(ctx *gin.Context) {
	inventoryIDParam := ctx.Param("id")
	inventoryID, err := uuid.Parse(inventoryIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.ClinicInventoryUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdinventory, err := h.clinicInventoryUsecase.Update(ctx, inventoryID, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create clinic inventory", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic inventory created successfully", createdinventory)
	ctx.JSON(http.StatusCreated, res)
}

func (h *ClinicInventoryHandler) Delete(ctx *gin.Context) {
	inventoryIDParam := ctx.Param("id")
	inventoryID, err := uuid.Parse(inventoryIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid clinic ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.clinicInventoryUsecase.Delete(ctx, inventoryID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete clinic inventory", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Clinic inventory deleted successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}
