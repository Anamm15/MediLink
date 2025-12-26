package dto

import (
	"time"

	"MediLink/internal/domain/entity"

	"github.com/google/uuid"
)

type MedicalRecordResponse struct {
	ID          uuid.UUID `json:"id"`
	DoctorName  string    `json:"doctor_name"`
	PatientName string    `json:"patient_name"`
	Title       string    `json:"title"`
	Subjective  *string   `json:"subjective"`
	Objective   *string   `json:"objective"`
	Assessment  *string   `json:"assessment"`
	Plan        *string   `json:"plan"`
	CreatedAt   time.Time `json:"created_at"`
}

type MedicalRecordCreateRequest struct {
	DoctorID      uuid.UUID  `json:"doctor_id"`
	PatientID     uuid.UUID  `json:"patient_id"`
	AppointmentID *uuid.UUID `json:"appointment_id"`
	Title         string     `json:"title"`
	Subjective    *string    `json:"subjective"`
	Objective     *string    `json:"objective"`
	Assessment    *string    `json:"assessment"`
	Plan          *string    `json:"plan"`
}

type MedicalRecordUpdateRequest struct {
	Title      string  `json:"title"`
	Subjective *string `json:"subjective"`
	Objective  *string `json:"objective"`
	Assessment *string `json:"assessment"`
	Plan       *string `json:"plan"`
}

type MedicalRecordDeleteRequest struct {
	DoctorID uuid.UUID `json:"doctor_id"`
}

func ToMedicalRecordResponse(medicalRecord *entity.MedicalRecord) MedicalRecordResponse {
	return MedicalRecordResponse{
		ID:          medicalRecord.ID,
		DoctorName:  medicalRecord.Doctor.User.Name,
		PatientName: medicalRecord.Patient.User.Name,
		Title:       medicalRecord.Title,
		Subjective:  medicalRecord.Subjective,
		Objective:   medicalRecord.Objective,
		Assessment:  medicalRecord.Assessment,
		Plan:        medicalRecord.Plan,
		CreatedAt:   medicalRecord.CreatedAt,
	}
}

func ToListMedicalRecordResponse(medicalRecords []entity.MedicalRecord) []MedicalRecordResponse {
	var medicalRecordResponses []MedicalRecordResponse
	for _, medicalRecord := range medicalRecords {
		medicalRecordResponses = append(medicalRecordResponses, ToMedicalRecordResponse(&medicalRecord))
	}
	return medicalRecordResponses
}

func (medicalRecord *MedicalRecordCreateRequest) ToModel(entity *entity.MedicalRecord) {
	entity.DoctorID = medicalRecord.DoctorID
	entity.PatientID = medicalRecord.PatientID
	entity.AppointmentID = medicalRecord.AppointmentID
	entity.Title = medicalRecord.Title
	entity.Subjective = medicalRecord.Subjective
	entity.Objective = medicalRecord.Objective
	entity.Assessment = medicalRecord.Assessment
	entity.Plan = medicalRecord.Plan
}

func (medicalRecord *MedicalRecordUpdateRequest) ToModel(entity *entity.MedicalRecord) {
	if medicalRecord.Title != "" {
		entity.Title = medicalRecord.Title
	}
	if medicalRecord.Subjective != nil {
		entity.Subjective = medicalRecord.Subjective
	}
	if medicalRecord.Objective != nil {
		entity.Objective = medicalRecord.Objective
	}
	if medicalRecord.Assessment != nil {
		entity.Assessment = medicalRecord.Assessment
	}
	if medicalRecord.Plan != nil {
		entity.Plan = medicalRecord.Plan
	}
}
