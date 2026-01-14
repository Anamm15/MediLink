package dto

import (
	"time"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type MedicineItem struct {
	MedicineResponse
	Quantity int `json:"quantity"`
}

type PrescriptionMedicinesCreate struct {
	MedicineID uuid.UUID `json:"medicine_id" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required"`
}

type PrescriptionResponse struct {
	ID              uuid.UUID              `json:"id"`
	Patient         PatientMinimumResponse `json:"patient"`
	Doctor          DoctorMinimumResponse  `json:"doctor"`
	MedicalRecordID uuid.UUID              `json:"medical_record_id"`

	Notes      *string   `json:"notes,omitempty"`
	IsRedeemed bool      `json:"is_redeemed"`
	CreatedAt  time.Time `json:"created_at"`

	Medicines []MedicineItem `json:"medicines"`
}

type PrescriptionCreate struct {
	PatientID       uuid.UUID                     `json:"patient_id" binding:"required"`
	MedicalRecordID uuid.UUID                     `json:"medical_record_id" binding:"required"`
	Notes           *string                       `json:"notes,omitempty"`
	Medicines       []PrescriptionMedicinesCreate `json:"medicines" binding:"required"`
}

type PrescriptionUpdate struct {
	Notes      *string `json:"notes,omitempty"`
	IsRedeemed *bool   `json:"is_redeemed,omitempty"`
}

func ToPrescriptionResponse(prescription *entity.Prescription) *PrescriptionResponse {
	var doctor *DoctorMinimumResponse
	var patient *PatientMinimumResponse
	var items []MedicineItem

	doctor = &DoctorMinimumResponse{
		ID:             prescription.Doctor.ID,
		Name:           prescription.Doctor.User.Name,
		Specialization: prescription.Doctor.Specialization,
	}

	patient = &PatientMinimumResponse{
		ID:          prescription.Patient.ID,
		Name:        prescription.Patient.User.Name,
		Email:       prescription.Patient.User.Email,
		PhoneNumber: prescription.Patient.User.PhoneNumber,
	}

	for _, medicine := range prescription.Medicines {
		medicineItem := MedicineItem{
			MedicineResponse: *ToMedicineResponse(&medicine.Medicine),
			Quantity:         medicine.Quantity,
		}
		items = append(items, medicineItem)
	}

	return &PrescriptionResponse{
		ID:              prescription.ID,
		Patient:         *patient,
		Doctor:          *doctor,
		MedicalRecordID: prescription.MedicalRecordID,
		Notes:           prescription.Notes,
		IsRedeemed:      prescription.IsRedeemed,
		CreatedAt:       prescription.CreatedAt,
		Medicines:       items,
	}
}

func ToListPrescriptionResponseDTO(prescriptions []entity.Prescription) []PrescriptionResponse {
	var prescriptionsDTO []PrescriptionResponse
	for _, prescription := range prescriptions {
		prescriptionsDTO = append(prescriptionsDTO, *ToPrescriptionResponse(&prescription))
	}
	return prescriptionsDTO
}

func (dto *PrescriptionCreate) ToModel(prescription *entity.Prescription) {
	prescription.PatientID = dto.PatientID
	prescription.Notes = dto.Notes
	prescription.MedicalRecordID = dto.MedicalRecordID

	for _, medicine := range dto.Medicines {
		prescription.Medicines = append(prescription.Medicines, entity.PrescriptionItem{
			MedicineID: medicine.MedicineID,
			Quantity:   medicine.Quantity,
		})
	}
}

func (dto *PrescriptionUpdate) ToModel(prescription *entity.Prescription) *entity.Prescription {
	if dto.Notes != nil {
		prescription.Notes = dto.Notes
	}
	if dto.IsRedeemed != nil {
		prescription.IsRedeemed = *dto.IsRedeemed
	}
	return prescription
}
