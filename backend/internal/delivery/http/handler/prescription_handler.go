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

type PrescriptionHandler struct {
	prescriptionUsecase usecase.PrescriptionUsecase
}

func NewPrescriptionHandler(prescriptionUsecase usecase.PrescriptionUsecase) handler.PrescriptionHandler {
	return &PrescriptionHandler{prescriptionUsecase: prescriptionUsecase}
}

func (ph *PrescriptionHandler) GetByPatient(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	prescriptions, err := ph.prescriptionUsecase.GetByPatient(c.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve prescriptions", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Prescriptions retrieved successfully", prescriptions)
	c.JSON(http.StatusOK, res)
}

func (ph *PrescriptionHandler) GetByDoctor(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	prescriptions, err := ph.prescriptionUsecase.GetByDoctor(c.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve prescriptions", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Prescriptions retrieved successfully", prescriptions)
	c.JSON(http.StatusOK, res)
}

func (ph *PrescriptionHandler) GetDetailByID(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid prescription ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	prescription, err := ph.prescriptionUsecase.GetDetailByID(c.Request.Context(), parsedID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve prescription", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Prescription retrieved successfully", prescription)
	c.JSON(http.StatusOK, res)
}

func (ph *PrescriptionHandler) Create(c *gin.Context) {
	var req dto.PrescriptionCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	prescription, err := ph.prescriptionUsecase.Create(c.Request.Context(), &req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create prescription", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("Prescription created successfully", prescription)
	c.JSON(http.StatusCreated, res)
}

func (ph *PrescriptionHandler) Update(c *gin.Context) {
	prescriptionIDParam := c.Param("id")
	prescriptionID, err := uuid.Parse(prescriptionIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid prescription ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.PrescriptionUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	updatedPrescription, err := ph.prescriptionUsecase.Update(c.Request.Context(), prescriptionID, &req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update prescription", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Prescription updated successfully", updatedPrescription)
	c.JSON(http.StatusOK, res)
}

func (ph *PrescriptionHandler) Delete(c *gin.Context) {
	prescriptionIDParam := c.Param("id")
	prescriptionID, err := uuid.Parse(prescriptionIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid prescription ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = ph.prescriptionUsecase.Delete(c.Request.Context(), prescriptionID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete prescription", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Prescription deleted successfully", nil)
	c.JSON(http.StatusNoContent, res)
}

func (ph *PrescriptionHandler) AddMedicine(c *gin.Context) {
	prescriptionIDParam := c.Param("id")
	prescriptionID, err := uuid.Parse(prescriptionIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid prescription ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.PrescriptionItemCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	medicine, err := ph.prescriptionUsecase.AddMedicine(c.Request.Context(), prescriptionID, &req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to add medicine to prescription", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicine added to prescription successfully", medicine)
	c.JSON(http.StatusCreated, res)
}

func (ph *PrescriptionHandler) UpdateMedicine(c *gin.Context) {
	prescriptionMedicineIDParam := c.Param("prescription_medicine_id")
	prescriptionMedicineID, err := uuid.Parse(prescriptionMedicineIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid prescription medicine ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.PrescriptionItemUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = ph.prescriptionUsecase.UpdateMedicine(c.Request.Context(), prescriptionMedicineID, req.Quantity)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update medicine in prescription", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicine updated in prescription successfully", nil)
	c.JSON(http.StatusOK, res)
}

func (ph *PrescriptionHandler) RemoveMedicine(c *gin.Context) {
	prescriptionMedicineIDParam := c.Param("prescription_medicine_id")
	prescriptionMedicineID, err := uuid.Parse(prescriptionMedicineIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid prescription medicine ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = ph.prescriptionUsecase.RemoveMedicine(c.Request.Context(), prescriptionMedicineID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to remove medicine from prescription", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medicine removed from prescription successfully", nil)
	c.JSON(http.StatusNoContent, res)
}
