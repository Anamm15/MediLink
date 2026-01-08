package utils

import (
	"MediLink/internal/helpers/enum"
)

func MapMidtransStatus(status, fraud string) (enum.PaymentStatus, enum.AppointmentStatus) {
	switch status {
	case "capture":
		if fraud == "challenge" {
			return enum.PaymentPending, ""
		} else if fraud == "accept" {
			return enum.PaymentPaid, enum.AppointmentConfirmed
		}
	case "settlement":
		return enum.PaymentPaid, enum.AppointmentConfirmed
	case "deny", "cancel", "expire":
		return enum.PaymentFailed, enum.AppointmentCanceled
	case "pending":
		return enum.PaymentPending, ""
	}

	return enum.PaymentPending, ""
}
