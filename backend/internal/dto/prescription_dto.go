package dto

import (
	"time"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type MedicineItem struct {
	MedicineResponseDTO
	Quantity int `json:"quantity"`
}

type PrescriptionMedicinesCreate struct {
	MedicineID uuid.UUID `json:"medicine_id" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required"`
}

type PrescriptionResponseDTO struct {
	ID              uuid.UUID `json:"id"`
	PatientID       uuid.UUID `json:"patient_id"`
	DoctorID        uuid.UUID `json:"doctor_id"`
	ClinicID        uuid.UUID `json:"clinic_id,omitempty"`
	MedicalRecordID uuid.UUID `json:"medical_record_id"`

	Notes      *string   `json:"notes,omitempty"`
	IsRedeemed bool      `json:"is_redeemed"`
	CreatedAt  time.Time `json:"created_at"`

	PatientName     *string        `json:"patient_name,omitempty"`
	DoctorName      *string        `json:"doctor_name,omitempty"`
	DoctorSpecialty *string        `json:"doctor_specialty,omitempty"`
	Medicines       []MedicineItem `json:"medicines"`
}

type PrescriptionCreateDTO struct {
	PatientID       uuid.UUID                     `json:"patient_id" binding:"required"`
	DoctorID        uuid.UUID                     `json:"doctor_id" binding:"required"`
	ClinicID        uuid.UUID                     `json:"clinic_id,omitempty"`
	MedicalRecordID uuid.UUID                     `json:"medical_record_id" binding:"required"`
	Notes           *string                       `json:"notes,omitempty"`
	Medicines       []PrescriptionMedicinesCreate `json:"medicines" binding:"required"`
}

type PrescriptionUpdateDTO struct {
	Notes      *string `json:"notes,omitempty"`
	IsRedeemed *bool   `json:"is_redeemed,omitempty"`
}

func ToPrescriptionResponseDTO(prescription *entity.Prescription) *PrescriptionResponseDTO {
	var doctorName *string
	var doctorSpecialty *string
	var patientName *string
	var items []MedicineItem
	if prescription.Doctor != nil {
		fullName := prescription.Doctor.User.Name
		doctorName = &fullName

		specialty := prescription.Doctor.Specialization
		doctorSpecialty = &specialty
	}

	if prescription.Patient != nil {
		fullName := prescription.Patient.User.Name
		patientName = &fullName
	}

	for _, medicine := range prescription.Medicines {
		medicineItem := MedicineItem{
			MedicineResponseDTO: *ToMedicineResponseDTO(&medicine.Medicine),
			Quantity:            medicine.Quantity,
		}
		items = append(items, medicineItem)
	}

	return &PrescriptionResponseDTO{
		ID:              prescription.ID,
		PatientID:       prescription.PatientID,
		DoctorID:        prescription.DoctorID,
		MedicalRecordID: prescription.MedicalRecordID,
		ClinicID:        prescription.ClinicID,
		Notes:           prescription.Notes,
		IsRedeemed:      prescription.IsRedeemed,
		CreatedAt:       prescription.CreatedAt,
		DoctorName:      doctorName,
		DoctorSpecialty: doctorSpecialty,
		PatientName:     patientName,
		Medicines:       items,
	}
}

func ToPrescriptionListResponseDTO(prescriptions []entity.Prescription) []PrescriptionResponseDTO {
	var prescriptionsDTO []PrescriptionResponseDTO
	for _, prescription := range prescriptions {
		prescriptionsDTO = append(prescriptionsDTO, *ToPrescriptionResponseDTO(&prescription))
	}
	return prescriptionsDTO
}

func (dto *PrescriptionCreateDTO) ToModel(prescription *entity.Prescription) {
	prescription.PatientID = dto.PatientID
	prescription.DoctorID = dto.DoctorID
	prescription.ClinicID = dto.ClinicID
	prescription.Notes = dto.Notes
	prescription.MedicalRecordID = dto.MedicalRecordID

	for _, medicine := range dto.Medicines {
		prescription.Medicines = append(prescription.Medicines, entity.PrescriptionItem{
			MedicineID: medicine.MedicineID,
			Quantity:   medicine.Quantity,
		})
	}
}

func (dto *PrescriptionUpdateDTO) ToModel(prescription *entity.Prescription) *entity.Prescription {
	if dto.Notes != nil {
		prescription.Notes = dto.Notes
	}
	if dto.IsRedeemed != nil {
		prescription.IsRedeemed = *dto.IsRedeemed
	}
	return prescription
}
