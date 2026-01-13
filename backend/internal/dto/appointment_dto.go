package dto

import (
	"MediLink/internal/domain/entity"
	"MediLink/internal/helpers/enum"
	"MediLink/internal/utils"

	"github.com/google/uuid"
)

type AppointmentDetailResponse struct {
	ID              uuid.UUID              `json:"id"`
	Doctor          DoctorMinimumResponse  `json:"doctor"`
	Patient         PatientMinimumResponse `json:"patient"`
	AppointmentDate string                 `json:"appointment_date"`
	StartTime       string                 `json:"start_time"`
	EndTime         string                 `json:"end_time"`
	Type            enum.AppointmentType   `json:"type"`
	Status          enum.AppointmentStatus `json:"status"`

	ConsultationFeeSnapshot float64 `json:"consultation_fee_snapshot"`
	QueueNumber             *int    `json:"queue_number"`
	MeetingLink             *string `json:"meeting_link"`

	SymptomComplaint *string `json:"symptom_complaint"`
	DoctorNotes      *string `json:"doctor_notes"`
}

type CreateBookingRequest struct {
	DoctorID         uuid.UUID `json:"doctor_id" binding:"required"`
	ScheduleID       uuid.UUID `json:"schedule_id" binding:"required"`
	AppointmentDate  string    `json:"appointment_date" binding:"required"`
	SymptomComplaint *string   `json:"symptom_complaint"`
}

func ToAppointmentDetailResponse(appointment *entity.Appointment) *AppointmentDetailResponse {
	doctorResponse := DoctorMinimumResponse{
		ID:             appointment.Doctor.User.ID,
		Name:           appointment.Doctor.User.Name,
		Specialization: appointment.Doctor.Specialization,
		PhoneNumber:    appointment.Doctor.User.PhoneNumber,
	}

	patientResponse := PatientMinimumResponse{
		ID:          appointment.Patient.User.ID,
		Name:        appointment.Patient.User.Name,
		Email:       appointment.Patient.User.Email,
		PhoneNumber: appointment.Patient.User.PhoneNumber,
	}

	return &AppointmentDetailResponse{
		ID:                      appointment.ID,
		Doctor:                  doctorResponse,
		Patient:                 patientResponse,
		AppointmentDate:         utils.FormatDate(appointment.AppointmentDate),
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

func (dto *CreateBookingRequest) ToModel(appointment *entity.Appointment) {
	appointment.DoctorID = dto.DoctorID
	appointment.AppointmentDate = utils.ParseDate(dto.AppointmentDate)
	appointment.SymptomComplaint = dto.SymptomComplaint
}
