package dto

import (
	"time"

	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type AppointmentDetailResponse struct {
	ID                   uuid.UUID              `json:"id"`
	DoctorName           string                 `json:"doctor_name"`
	DoctorSpecialization string                 `json:"doctor_specialization"`
	DoctorPhoneNumber    string                 `json:"doctor_phone_number"`
	PatientName          string                 `json:"patient_name"`
	AppointmentDate      time.Time              `json:"appointment_date"`
	StartTime            time.Time              `json:"start_time"`
	EndTime              time.Time              `json:"end_time"`
	Type                 enum.AppointmentType   `json:"type"`
	Status               enum.AppointmentStatus `json:"status"`

	ConsultationFeeSnapshot float64 `json:"consultation_fee_snapshot"`
	QueueNumber             *int    `json:"queue_number"`
	MeetingLink             *string `json:"meeting_link"`

	SymptomComplaint *string `json:"symptom_complaint"`
	DoctorNotes      *string `json:"doctor_notes"`
}

type AppointmentCreateRequest struct {
	DoctorID        uuid.UUID            `json:"doctor_id" binding:"required"`
	PatientID       uuid.UUID            `json:"patient_id" binding:"required"`
	ClinicID        uuid.UUID            `json:"clinic_id" binding:"required"`
	AppointmentDate time.Time            `json:"appointment_date" binding:"required"`
	StartTime       time.Time            `json:"start_time" binding:"required"`
	EndTime         time.Time            `json:"end_time" binding:"required"`
	Type            enum.AppointmentType `json:"type" binding:"required"`

	ConsultationFeeSnapshot float64 `json:"consultation_fee_snapshot" binding:"required"`
	QueueNumber             *int    `json:"queue_number"`
	MeetingLink             *string `json:"meeting_link"`

	SymptomComplaint *string `json:"symptom_complaint"`
	DoctorNotes      *string `json:"doctor_notes"`
}

func ToAppointmentDetailResponse(appointment *entity.Appointment) *AppointmentDetailResponse {
	return &AppointmentDetailResponse{
		ID:                      appointment.ID,
		DoctorName:              appointment.Doctor.User.Name,
		DoctorSpecialization:    appointment.Doctor.Specialization,
		DoctorPhoneNumber:       appointment.Doctor.User.PhoneNumber,
		PatientName:             appointment.Patient.User.Name,
		AppointmentDate:         appointment.AppointmentDate,
		StartTime:               appointment.StartTime,
		EndTime:                 appointment.EndTime,
		Type:                    appointment.Type,
		Status:                  appointment.Status,
		ConsultationFeeSnapshot: appointment.ConsultationFeeSnapshot,
		QueueNumber:             appointment.QueueNumber,
		MeetingLink:             appointment.MeetingLink,
		SymptomComplaint:        appointment.SymptomComplaint,
		DoctorNotes:             appointment.DoctorNotes,
	}
}

func ToListAppointmentDetailResponse(appointments []entity.Appointment) []AppointmentDetailResponse {
	var appointmentResponses []AppointmentDetailResponse
	for _, appointment := range appointments {
		appointmentResponses = append(appointmentResponses, *ToAppointmentDetailResponse(&appointment))
	}
	return appointmentResponses
}

func (r *AppointmentCreateRequest) ToModel(appointment *entity.Appointment) {
	appointment.DoctorID = r.DoctorID
	appointment.PatientID = r.PatientID
	appointment.ClinicID = r.ClinicID
	appointment.AppointmentDate = r.AppointmentDate
	appointment.StartTime = r.StartTime
	appointment.EndTime = r.EndTime
	appointment.Type = r.Type

	appointment.ConsultationFeeSnapshot = r.ConsultationFeeSnapshot
	appointment.QueueNumber = r.QueueNumber
	appointment.MeetingLink = r.MeetingLink

	appointment.SymptomComplaint = r.SymptomComplaint
	appointment.DoctorNotes = r.DoctorNotes
}
