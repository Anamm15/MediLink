package dto

import (
	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PrescriptionMedicineResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	PrescriptionID uuid.UUID `json:"prescription_id"`
	MedicineID     uuid.UUID `json:"medicine_id"`
	Quantity       int       `json:"quantity"`
}

type PrescriptionMedicineCreateDTO struct {
	PrescriptionID uuid.UUID `json:"prescription_id" binding:"required"`
	MedicineID     uuid.UUID `json:"medicine_id" binding:"required"`
	Quantity       int       `json:"quantity" binding:"required"`
}

func (dto *PrescriptionMedicineCreateDTO) ToModel(pm *entity.PrescriptionMedicine) {
	pm.PrescriptionID = dto.PrescriptionID
	pm.MedicineID = dto.MedicineID
	pm.Quantity = dto.Quantity
}

type PrescriptionMedicineUpdateDTO struct {
	Quantity int `json:"quantity" binding:"required"`
}

func ToPrescriptionMedicineResponseDTO(data *entity.PrescriptionMedicine) PrescriptionMedicineResponseDTO {
	return PrescriptionMedicineResponseDTO{
		ID:             data.ID,
		PrescriptionID: data.PrescriptionID,
		MedicineID:     data.MedicineID,
		Quantity:       data.Quantity,
	}
}
