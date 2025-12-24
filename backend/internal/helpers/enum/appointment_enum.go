package enum

type (
	AppointmentStatus string
	AppointmentType   string
)

const (
	AppointmentPending    AppointmentStatus = "pending"
	AppointmentConfirmed  AppointmentStatus = "confirmed"
	AppointmentCompleted  AppointmentStatus = "completed"
	AppointmentInProgress AppointmentStatus = "in_progress"
	AppointmentCanceled   AppointmentStatus = "canceled"
	AppointmentExpired    AppointmentStatus = "expired"
)

const (
	AppointmentVideoCall AppointmentType = "video_call"
	AppointmentChat      AppointmentType = "chat"
	AppointmentOnsite    AppointmentType = "onsite"
)
