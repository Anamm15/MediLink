package dto

import (
	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type MedicineResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	GenericName  *string   `json:"generic_name,omitempty"`
	Description  *string   `json:"description"`
	Category     *string   `json:"category,omitempty"`
	Manufacturer *string   `json:"manufacturer,omitempty"`

	BasePrice              float64 `json:"base_price"`
	IsPrescriptionRequired bool    `json:"is_prescription_required"`

	CreatedAt string `json:"created_at"`
}

type MedicineCreate struct {
	Name         string  `json:"name" binding:"required"`
	Description  *string `json:"description"`
	GenericName  *string `json:"generic_name"`
	Category     *string `json:"category"`
	Manufacturer *string `json:"manufacturer"`

	BasePrice              float64 `json:"base_price" binding:"required"`
	IsPrescriptionRequired bool    `json:"is_prescription_required"`
}

type MedicineUpdate struct {
	Name         *string `json:"name"`
	GenericName  *string `json:"generic_name"`
	Description  *string `json:"description"`
	Category     *string `json:"category"`
	Manufacturer *string `json:"manufacturer"`

	BasePrice              *float64 `json:"base_price"`
	IsPrescriptionRequired *bool    `json:"is_prescription_required"`
}

func (m *MedicineCreate) ToModel() *entity.Medicine {
	return &entity.Medicine{
		Name:                   m.Name,
		Description:            m.Description,
		GenericName:            m.GenericName,
		Category:               m.Category,
		Manufacturer:           m.Manufacturer,
		BasePrice:              m.BasePrice,
		IsPrescriptionRequired: m.IsPrescriptionRequired,
	}
}

func (m *MedicineUpdate) ToModel(existing *entity.Medicine) *entity.Medicine {
	if m.Name != nil {
		existing.Name = *m.Name
	}
	if m.Description != nil {
		existing.Description = m.Description
	}
	if m.GenericName != nil {
		existing.GenericName = m.GenericName
	}
	if m.Category != nil {
		existing.Category = m.Category
	}
	if m.Manufacturer != nil {
		existing.Manufacturer = m.Manufacturer
	}
	if m.BasePrice != nil {
		existing.BasePrice = *m.BasePrice
	}
	if m.IsPrescriptionRequired != nil {
		existing.IsPrescriptionRequired = *m.IsPrescriptionRequired
	}
	return existing
}

func ToMedicineResponse(m *entity.Medicine) *MedicineResponse {
	return &MedicineResponse{
		ID:                     m.ID,
		Name:                   m.Name,
		Description:            m.Description,
		GenericName:            m.GenericName,
		Category:               m.Category,
		Manufacturer:           m.Manufacturer,
		BasePrice:              m.BasePrice,
		IsPrescriptionRequired: m.IsPrescriptionRequired,
	}
}

func ToListMedicineResponse(medicines []entity.Medicine) []MedicineResponse {
	response := make([]MedicineResponse, 0, len(medicines))
	for _, m := range medicines {
		response = append(response, *ToMedicineResponse(&m))
	}
	return response
}
