package enum

type (
	AppointmentStatus     string
	AppointmentType       string
	AppointmentDateStatus string
)

const (
	AppointmentPending    AppointmentStatus = "pending"
	AppointmentConfirmed  AppointmentStatus = "confirmed"
	AppointmentCompleted  AppointmentStatus = "completed"
	AppointmentInProgress AppointmentStatus = "in_progress"
	AppointmentCanceled   AppointmentStatus = "canceled"
)

const (
	AppointmentVideoCall AppointmentType = "video_call"
	AppointmentChat      AppointmentType = "chat"
	AppointmentOnsite    AppointmentType = "onsite"
)

const (
	AppointmentUpcoming AppointmentDateStatus = "upcoming"
	AppointmentPast     AppointmentDateStatus = "past"
)
