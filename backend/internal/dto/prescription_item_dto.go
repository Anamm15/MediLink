package dto

import (
	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PrescriptionItemResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	PrescriptionID uuid.UUID `json:"prescription_id"`
	MedicineID     uuid.UUID `json:"medicine_id"`

	Quantity     int     `json:"quantity"`
	Instructions *string `json:"instructions,omitempty"`
}

type PrescriptionItemCreateDTO struct {
	PrescriptionID uuid.UUID `json:"prescription_id" binding:"required"`
	MedicineID     uuid.UUID `json:"medicine_id" binding:"required"`
	Quantity       int       `json:"quantity" binding:"required"`
	Instructions   *string   `json:"instructions,omitempty"`
}

func (dto *PrescriptionItemCreateDTO) ToModel(pm *entity.PrescriptionItem) {
	pm.PrescriptionID = dto.PrescriptionID
	pm.MedicineID = dto.MedicineID
	pm.Quantity = dto.Quantity
	pm.Instructions = dto.Instructions
}

type PrescriptionItemUpdateDTO struct {
	Quantity     int     `json:"quantity" binding:"required"`
	Instructions *string `json:"instructions,omitempty"`
}

func ToPrescriptionItemResponseDTO(data *entity.PrescriptionItem) PrescriptionItemResponseDTO {
	return PrescriptionItemResponseDTO{
		ID:             data.ID,
		PrescriptionID: data.PrescriptionID,
		MedicineID:     data.MedicineID,
		Quantity:       data.Quantity,
		Instructions:   data.Instructions,
	}
}
