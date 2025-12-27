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

type MedicineHandler struct {
	medicineUsecase usecase.MedicineUsecase
}

func NewMedicineHandler(medicineUsecase usecase.MedicineUsecase) handler.MedicineHandler {
	return &MedicineHandler{medicineUsecase: medicineUsecase}
}

func (h *MedicineHandler) GetAll(ctx *gin.Context) {
	page := ctx.Query("page")
	medicines, err := h.medicineUsecase.GetAll(ctx.Request.Context(), utils.StringToInt(page))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve medicines", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicines retrieved successfully", medicines)
	ctx.JSON(http.StatusOK, res)
}

func (h *MedicineHandler) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	medicineID, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid medicine ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	medicine, err := h.medicineUsecase.GetByID(ctx.Request.Context(), medicineID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve medicine", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicine retrieved successfully", medicine)
	ctx.JSON(http.StatusOK, res)
}

func (h *MedicineHandler) Search(ctx *gin.Context) {
	name := ctx.Query("name")
	page := ctx.Query("page")
	medicines, err := h.medicineUsecase.Search(ctx.Request.Context(), name, utils.StringToInt(page))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to search medicines", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicines found successfully", medicines)
	ctx.JSON(http.StatusOK, res)
}

func (h *MedicineHandler) Create(ctx *gin.Context) {
	var req dto.MedicineCreate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
	}

	createdMedicine, err := h.medicineUsecase.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create medicine", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicine created successfully", createdMedicine)
	ctx.JSON(http.StatusCreated, res)
}

func (h *MedicineHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	medicineID, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid medicine ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.MedicineUpdate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	updatedMedicine, err := h.medicineUsecase.Update(ctx.Request.Context(), medicineID, &req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update medicine", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicine updated successfully", updatedMedicine)
	ctx.JSON(http.StatusOK, res)
}

func (h *MedicineHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	medicineID, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid medicine ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.medicineUsecase.Delete(ctx.Request.Context(), medicineID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete medicine", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicine deleted successfully", nil)
	ctx.JSON(http.StatusNoContent, res)
}
