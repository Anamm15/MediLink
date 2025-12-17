package constants

type (
	AppointmentStatus string
	AppointmentType   string
	ScheduleDay       string
)

const (
	AppointmentPending   AppointmentStatus = "pending"
	AppointmentConfirmed AppointmentStatus = "confirmed"
	AppointmentCompleted AppointmentStatus = "completed"
	AppointmentCanceled  AppointmentStatus = "canceled"
	AppointmentExpired   AppointmentStatus = "expired"
)

const (
	AppointmentVideoCall AppointmentType = "video_call"
	AppointmentChat      AppointmentType = "chat"
	AppointmentOnsite    AppointmentType = "onsite"
)

const (
	DayMonday    ScheduleDay = "monday"
	DayTuesday   ScheduleDay = "tuesday"
	DayWednesday ScheduleDay = "wednesday"
	DayThursday  ScheduleDay = "thursday"
	DayFriday    ScheduleDay = "friday"
	DaySaturday  ScheduleDay = "saturday"
	DaySunday    ScheduleDay = "sunday"
)
