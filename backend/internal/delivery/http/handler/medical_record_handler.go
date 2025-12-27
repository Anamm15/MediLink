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

type MedicalRecordHandler struct {
	medicalRecordUsecase usecase.MedicalRecordUsecase
}

func NewMedicalRecordHandler(medicalRecordUsecase usecase.MedicalRecordUsecase) handler.MedicalRecordHandler {
	return &MedicalRecordHandler{medicalRecordUsecase: medicalRecordUsecase}
}

func (h *MedicalRecordHandler) GetByPatient(c *gin.Context) {
	patientIDParam := c.Param("patient_id")
	patientID, err := uuid.Parse(patientIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid patient ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	medicalRecords, err := h.medicalRecordUsecase.GetByPatient(c.Request.Context(), patientID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve medical records", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medical records retrieved successfully", medicalRecords)
	c.JSON(http.StatusOK, res)
}

func (h *MedicalRecordHandler) GetByDoctor(c *gin.Context) {
	doctorIDParam := c.Param("doctor_id")
	doctorID, err := uuid.Parse(doctorIDParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid patient ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	medicalRecords, err := h.medicalRecordUsecase.GetByDoctor(c.Request.Context(), doctorID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve medical records", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medical records retrieved successfully", medicalRecords)
	c.JSON(http.StatusOK, res)
}

func (h *MedicalRecordHandler) GetDetailByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid patient ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	medicalRecords, err := h.medicalRecordUsecase.GetById(c.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve medical records", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medical records retrieved successfully", medicalRecords)
	c.JSON(http.StatusOK, res)
}

func (h *MedicalRecordHandler) Create(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	var req dto.MedicalRecordCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	medicalRecord, err := h.medicalRecordUsecase.Create(c.Request.Context(), userID, &req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create medical record", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medical record created successfully", medicalRecord)
	c.JSON(http.StatusCreated, res)
}

func (h *MedicalRecordHandler) Update(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid patient ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.MedicalRecordUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	medicalRecord, err := h.medicalRecordUsecase.Update(c.Request.Context(), id, userID, &req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update medical record", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medical record updated successfully", medicalRecord)
	c.JSON(http.StatusOK, res)
}

func (h *MedicalRecordHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid patient ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.MedicalRecordDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.medicalRecordUsecase.Delete(c.Request.Context(), id, req.DoctorID)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete medical record", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Medical record deleted successfully", nil)
	c.JSON(http.StatusNoContent, res)
}
