package enum

type (
	UserRole   string
	UserStatus string
	Gender     string
)

const (
	RoleUser       UserRole = "user"
	RolePatient    UserRole = "patient"
	RoleDoctor     UserRole = "doctor"
	RoleAdmin      UserRole = "admin"
	RoleNurse      UserRole = "nurse"
	RolePharmacist UserRole = "pharmacist"
	RoleSuperAdmin UserRole = "super_admin"
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
