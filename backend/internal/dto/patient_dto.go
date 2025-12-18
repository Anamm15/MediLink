package dto

import (
	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type PatientResponseDTO struct {
	PatientID         uuid.UUID `json:"patient_id"`
	IdentityNumber    string    `json:"identity_number"`
	BloodType         *string   `json:"blood_type"`
	WeightKg          *float64  `json:"weight_kg"`
	HeightCm          *float64  `json:"height_cm"`
	Allergies         *string   `json:"allergies"`
	ChronicDiseases   *string   `json:"chronic_diseases"`
	EmergencyContact  *string   `json:"emergency_contact"`
	InsuranceProvider *string   `json:"insurance_provider"`
	InsuranceNumber   *string   `json:"insurance_number"`
	Occupation        *string   `json:"occupation"`
}

type PatientCreateRequestDTO struct {
	IdentityNumber    string   `json:"identity_number" binding:"required"`
	BloodType         *string  `json:"blood_type,omitempty"`
	WeightKg          *float64 `json:"weight_kg,omitempty"`
	HeightCm          *float64 `json:"height_cm,omitempty"`
	Allergies         *string  `json:"allergies,omitempty"`
	ChronicDiseases   *string  `json:"chronic_diseases,omitempty"`
	EmergencyContact  *string  `json:"emergency_contact,omitempty"`
	InsuranceProvider *string  `json:"insurance_provider,omitempty"`
	InsuranceNumber   *string  `json:"insurance_number,omitempty"`
	Occupation        *string  `json:"occupation,omitempty"`
}

type PatientUpdateRequestDTO struct {
	IdentityNumber    *string  `json:"identity_number"`
	BloodType         *string  `json:"blood_type,omitempty"`
	WeightKg          *float64 `json:"weight_kg,omitempty"`
	HeightCm          *float64 `json:"height_cm,omitempty"`
	Allergies         *string  `json:"allergies,omitempty"`
	ChronicDiseases   *string  `json:"chronic_diseases,omitempty"`
	EmergencyContact  *string  `json:"emergency_contact,omitempty"`
	InsuranceProvider *string  `json:"insurance_provider,omitempty"`
	InsuranceNumber   *string  `json:"insurance_number,omitempty"`
	Occupation        *string  `json:"occupation,omitempty"`
}

func MapPatientToPatientResponseDTO(patient *entity.Patient) PatientResponseDTO {
	return PatientResponseDTO{
		PatientID:         patient.ID,
		IdentityNumber:    *patient.IdentityNumber,
		BloodType:         patient.BloodType,
		WeightKg:          patient.WeightKg,
		HeightCm:          patient.HeightCm,
		Allergies:         patient.Allergies,
		ChronicDiseases:   patient.ChronicDiseases,
		EmergencyContact:  patient.EmergencyContact,
		InsuranceProvider: patient.InsuranceProvider,
		InsuranceNumber:   patient.InsuranceNumber,
		Occupation:        patient.Occupation,
	}
}

func (dto *PatientCreateRequestDTO) AssignToEntity(patient *entity.Patient) {
	patient.IdentityNumber = &dto.IdentityNumber
	patient.BloodType = dto.BloodType
	patient.WeightKg = dto.WeightKg
	patient.HeightCm = dto.HeightCm
	patient.Allergies = dto.Allergies
	patient.ChronicDiseases = dto.ChronicDiseases
	patient.EmergencyContact = dto.EmergencyContact
	patient.InsuranceProvider = dto.InsuranceProvider
	patient.InsuranceNumber = dto.InsuranceNumber
	patient.Occupation = dto.Occupation
}

func (dto *PatientUpdateRequestDTO) AssignToEntity(patient *entity.Patient) {
	if dto.IdentityNumber != nil {
		patient.IdentityNumber = dto.IdentityNumber
	}
	if dto.BloodType != nil {
		patient.BloodType = dto.BloodType
	}
	if dto.WeightKg != nil {
		patient.WeightKg = dto.WeightKg
	}
	if dto.HeightCm != nil {
		patient.HeightCm = dto.HeightCm
	}
	if dto.Allergies != nil {
		patient.Allergies = dto.Allergies
	}
	if dto.ChronicDiseases != nil {
		patient.ChronicDiseases = dto.ChronicDiseases
	}
	if dto.EmergencyContact != nil {
		patient.EmergencyContact = dto.EmergencyContact
	}
	if dto.InsuranceProvider != nil {
		patient.InsuranceProvider = dto.InsuranceProvider
	}
	if dto.InsuranceNumber != nil {
		patient.InsuranceNumber = dto.InsuranceNumber
	}
	if dto.Occupation != nil {
		patient.Occupation = dto.Occupation
	}
}
