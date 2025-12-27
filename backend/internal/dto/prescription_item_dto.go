package dto

import (
	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PrescriptionItemResponse struct {
	ID             uuid.UUID `json:"id"`
	PrescriptionID uuid.UUID `json:"prescription_id"`
	MedicineID     uuid.UUID `json:"medicine_id"`

	Quantity     int     `json:"quantity"`
	Instructions *string `json:"instructions,omitempty"`
}

type PrescriptionItemCreate struct {
	PrescriptionID uuid.UUID `json:"prescription_id" binding:"required"`
	MedicineID     uuid.UUID `json:"medicine_id" binding:"required"`
	Quantity       int       `json:"quantity" binding:"required"`
	Instructions   *string   `json:"instructions,omitempty"`
}

func (dto *PrescriptionItemCreate) ToModel(pm *entity.PrescriptionItem) {
	pm.PrescriptionID = dto.PrescriptionID
	pm.MedicineID = dto.MedicineID
	pm.Quantity = dto.Quantity
	pm.Instructions = dto.Instructions
}

type PrescriptionItemUpdate struct {
	Quantity     int     `json:"quantity" binding:"required"`
	Instructions *string `json:"instructions,omitempty"`
}

func ToPrescriptionItemResponse(data *entity.PrescriptionItem) PrescriptionItemResponse {
	return PrescriptionItemResponse{
		ID:             data.ID,
		PrescriptionID: data.PrescriptionID,
		MedicineID:     data.MedicineID,
		Quantity:       data.Quantity,
		Instructions:   data.Instructions,
	}
}
