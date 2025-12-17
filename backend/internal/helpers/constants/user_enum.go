package constants

type (
	UserRole   string
	UserStatus string
	Gender     string
)

const (
	RoleUser    UserRole = "user"
	RolePatient UserRole = "patient"
	RoleDoctor  UserRole = "doctor"
	RoleAdmin   UserRole = "admin"
	RoleStaff   UserRole = "staff"
	RoleClinic  UserRole = "clinic"
)

const (
	StatusActive    UserStatus = "active"
	StatusInactive  UserStatus = "inactive"
	StatusSuspended UserStatus = "suspended"
	StatusBanned    UserStatus = "banned"
)

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)
