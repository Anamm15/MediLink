package dto

import "github.com/google/uuid"

type BookingResponse struct {
	AppointmentID uuid.UUID `json:"appointment_id"`
	PaymentURL    *string   `json:"payment_url"`
}
