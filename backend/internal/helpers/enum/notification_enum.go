package enum

type OTPType string

const (
	OTPTypeRegister OTPType = "verify-user"
	OTPTypeLogin    OTPType = "change-password"
)
