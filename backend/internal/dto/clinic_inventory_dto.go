package dto

import (
	"time"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type ClinicInventoryResponse struct {
	ID         uuid.UUID `json:"id"`
	MedicineID uuid.UUID `json:"medicine_id"`
	ItemName   string    `json:"item_name"`
	ClinicName string    `json:"clinic_name"`

	CurrentStock      int     `json:"current_stock"`
	LowStockThreshold int     `json:"low_stock_threshold"`
	Price             float64 `json:"price"`

	BatchNumber *string `json:"batch_number"`
	ExpiryDate  *string `json:"expiry_date"`

	UpdatedAt time.Time `json:"updated_at"`
}

type ClinicInventoryCreateRequest struct {
	ClinicID   uuid.UUID `json:"clinic_id" binding:"required"`
	MedicineID uuid.UUID `json:"medicine_id" binding:"required"`

	CurrentStock      int     `json:"current_stock" binding:"required"`
	LowStockThreshold int     `json:"low_stock_threshold" binding:"required"`
	Price             float64 `json:"price" binding:"required"`

	BatchNumber *string `json:"batch_number"`
	ExpiryDate  *string `json:"expiry_date"`
}

type ClinicInventoryUpdateRequest struct {
	CurrentStock      *int     `json:"current_stock"`
	LowStockThreshold *int     `json:"low_stock_threshold"`
	Price             *float64 `json:"price"`
	BatchNumber       *string  `json:"batch_number"`
	ExpiryDate        *string  `json:"expiry_date"`
}

func ToClinicInventoryResponse(clinicInventory *entity.ClinicInventory) *ClinicInventoryResponse {
	return &ClinicInventoryResponse{
		ID:                clinicInventory.ID,
		MedicineID:        clinicInventory.MedicineID,
		ItemName:          clinicInventory.Medicine.Name,
		ClinicName:        clinicInventory.Clinic.Name,
		CurrentStock:      clinicInventory.CurrentStock,
		LowStockThreshold: clinicInventory.LowStockThreshold,
		Price:             clinicInventory.Price,
		BatchNumber:       clinicInventory.BatchNumber,
		ExpiryDate:        clinicInventory.ExpiryDate,
		UpdatedAt:         clinicInventory.UpdatedAt,
	}
}

func ToListClinicInventoryResponse(clinicInventories []entity.ClinicInventory) []ClinicInventoryResponse {
	var clinicInventoryResponse []ClinicInventoryResponse
	for _, clinicInventory := range clinicInventories {
		clinicInventoryResponse = append(clinicInventoryResponse, *ToClinicInventoryResponse(&clinicInventory))
	}
	return clinicInventoryResponse
}

func (dto *ClinicInventoryCreateRequest) ToModel(clinicInventory *entity.ClinicInventory) *entity.ClinicInventory {
	clinicInventory.ClinicID = dto.ClinicID
	clinicInventory.MedicineID = dto.MedicineID
	clinicInventory.CurrentStock = dto.CurrentStock
	clinicInventory.LowStockThreshold = dto.LowStockThreshold
	clinicInventory.Price = dto.Price
	clinicInventory.BatchNumber = dto.BatchNumber
	clinicInventory.ExpiryDate = dto.ExpiryDate
	return clinicInventory
}

func (dto *ClinicInventoryUpdateRequest) ToModel(clinicInventory *entity.ClinicInventory) *entity.ClinicInventory {
	if dto.CurrentStock != nil {
		clinicInventory.CurrentStock = *dto.CurrentStock
	}
	if dto.LowStockThreshold != nil {
		clinicInventory.LowStockThreshold = *dto.LowStockThreshold
	}
	if dto.Price != nil {
		clinicInventory.Price = *dto.Price
	}
	if dto.BatchNumber != nil {
		clinicInventory.BatchNumber = dto.BatchNumber
	}
	if dto.ExpiryDate != nil {
		clinicInventory.ExpiryDate = dto.ExpiryDate
	}
	return clinicInventory
}
