package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"
	"MediLink/internal/utils"

	"github.com/google/uuid"
)

type PatientResponseDTO struct {
	PatientID uuid.UUID `json:"patient_id"`

	BirthDate      time.Time   `json:"birth_date"`
	Gender         enum.Gender `json:"gender"`
	IdentityNumber string      `json:"identity_number"`
	BloodType      string      `json:"blood_type"`

	WeightKg               float64 `json:"weight_kg"`
	HeightCm               float64 `json:"height_cm"`
	Allergies              *string `json:"allergies"`
	HistoryChronicDiseases *string `json:"history_chronic_diseases"`
	EmergencyContact       *string `json:"emergency_contact"`

	InsuranceProvider *string `json:"insurance_provider"`
	InsuranceNumber   *string `json:"insurance_number"`
	Occupation        *string `json:"occupation"`
}

type PatientCreateRequestDTO struct {
	IdentityNumber string      `json:"identity_number" binding:"required"`
	BirthDate      string      `json:"birth_date" binding:"required"`
	Gender         enum.Gender `json:"gender" binding:"required"`
	BloodType      string      `json:"blood_type,omitempty"`

	WeightKg               float64 `json:"weight_kg,omitempty"`
	HeightCm               float64 `json:"height_cm,omitempty"`
	Allergies              *string `json:"allergies,omitempty"`
	HistoryChronicDiseases *string `json:"history_chronic_diseases,omitempty"`
	EmergencyContact       *string `json:"emergency_contact,omitempty"`

	InsuranceProvider *string `json:"insurance_provider,omitempty"`
	InsuranceNumber   *string `json:"insurance_number,omitempty"`
	Occupation        *string `json:"occupation,omitempty"`
}

type PatientUpdateRequestDTO struct {
	BirthDate *string      `json:"birth_date"`
	Gender    *enum.Gender `json:"gender"`
	BloodType *string      `json:"blood_type,omitempty"`

	WeightKg               *float64 `json:"weight_kg,omitempty"`
	HeightCm               *float64 `json:"height_cm,omitempty"`
	Allergies              *string  `json:"allergies,omitempty"`
	HistoryChronicDiseases *string  `json:"history_chronic_diseases,omitempty"`
	EmergencyContact       *string  `json:"emergency_contact,omitempty"`

	InsuranceProvider *string `json:"insurance_provider,omitempty"`
	InsuranceNumber   *string `json:"insurance_number,omitempty"`
	Occupation        *string `json:"occupation,omitempty"`
}

func MapPatientToPatientResponseDTO(patient *entity.Patient) PatientResponseDTO {
	return PatientResponseDTO{
		PatientID:              patient.ID,
		IdentityNumber:         patient.IdentityNumber,
		BirthDate:              patient.BirthDate,
		Gender:                 patient.Gender,
		BloodType:              patient.BloodType,
		WeightKg:               patient.WeightKg,
		HeightCm:               patient.HeightCm,
		Allergies:              patient.Allergies,
		HistoryChronicDiseases: patient.HistoryChronicDiseases,
		EmergencyContact:       patient.EmergencyContact,
		InsuranceProvider:      patient.InsuranceProvider,
		InsuranceNumber:        patient.InsuranceNumber,
		Occupation:             patient.Occupation,
	}
}

func (dto *PatientCreateRequestDTO) AssignToEntity(patient *entity.Patient) {
	patient.IdentityNumber = dto.IdentityNumber
	patient.BirthDate = *utils.ConvertStringToTime(dto.BirthDate)
	patient.Gender = dto.Gender
	patient.BloodType = dto.BloodType
	patient.WeightKg = dto.WeightKg
	patient.HeightCm = dto.HeightCm
	patient.Allergies = dto.Allergies
	patient.HistoryChronicDiseases = dto.HistoryChronicDiseases
	patient.EmergencyContact = dto.EmergencyContact
	patient.InsuranceProvider = dto.InsuranceProvider
	patient.InsuranceNumber = dto.InsuranceNumber
	patient.Occupation = dto.Occupation
}

func (dto *PatientUpdateRequestDTO) AssignToEntity(patient *entity.Patient) {
	if dto.BirthDate != nil {
		patient.BirthDate = *utils.ConvertStringToTime(*dto.BirthDate)
	}
	if dto.Gender != nil {
		patient.Gender = *dto.Gender
	}
	if dto.BloodType != nil {
		patient.BloodType = *dto.BloodType
	}
	if dto.WeightKg != nil {
		patient.WeightKg = *dto.WeightKg
	}
	if dto.HeightCm != nil {
		patient.HeightCm = *dto.HeightCm
	}
	if dto.Allergies != nil {
		patient.Allergies = dto.Allergies
	}
	if dto.HistoryChronicDiseases != nil {
		patient.HistoryChronicDiseases = dto.HistoryChronicDiseases
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
