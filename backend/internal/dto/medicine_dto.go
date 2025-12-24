package dto

import (
	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type MedicineResponseDTO struct {
	ID                   uuid.UUID `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Dosage               string    `json:"dosage"`
	Price                float64   `json:"price"`
	Stock                int       `json:"stock"`
	RequiresPrescription bool      `json:"requires_prescription"`
}

type MedicineCreateDTO struct {
	Name                 string  `json:"name" binding:"required"`
	Description          string  `json:"description" binding:"required"`
	Dosage               string  `json:"dosage" binding:"required"`
	Price                float64 `json:"price" binding:"required"`
	Stock                int     `json:"stock" binding:"required"`
	RequiresPrescription bool    `json:"requires_prescription"`
}

type MedicineUpdateDTO struct {
	Name                 *string  `json:"name"`
	Description          *string  `json:"description"`
	Dosage               *string  `json:"dosage"`
	Price                *float64 `json:"price"`
	Stock                *int     `json:"stock"`
	RequiresPrescription *bool    `json:"requires_prescription"`
}

func (m *MedicineCreateDTO) ToModel() *entity.Medicine {
	return &entity.Medicine{
		Name:                 m.Name,
		Description:          m.Description,
		Dosage:               m.Dosage,
		Price:                m.Price,
		Stock:                m.Stock,
		RequiresPrescription: m.RequiresPrescription,
	}
}

func (m *MedicineUpdateDTO) ToModel(existing *entity.Medicine) *entity.Medicine {
	if m.Name != nil {
		existing.Name = *m.Name
	}
	if m.Description != nil {
		existing.Description = *m.Description
	}
	if m.Dosage != nil {
		existing.Dosage = *m.Dosage
	}
	if m.Price != nil {
		existing.Price = *m.Price
	}
	if m.Stock != nil {
		existing.Stock = *m.Stock
	}
	if m.RequiresPrescription != nil {
		existing.RequiresPrescription = *m.RequiresPrescription
	}
	return existing
}

func ToMedicineResponseDTO(m *entity.Medicine) *MedicineResponseDTO {
	return &MedicineResponseDTO{
		ID:                   m.ID,
		Name:                 m.Name,
		Description:          m.Description,
		Dosage:               m.Dosage,
		Price:                m.Price,
		Stock:                m.Stock,
		RequiresPrescription: m.RequiresPrescription,
	}
}

func ToListMedicineResponseDTO(medicines []entity.Medicine) []MedicineResponseDTO {
	response := make([]MedicineResponseDTO, 0, len(medicines))
	for _, m := range medicines {
		response = append(response, *ToMedicineResponseDTO(&m))
	}
	return response
}
