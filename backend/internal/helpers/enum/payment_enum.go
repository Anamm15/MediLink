package enum

type PaymentStatus string

const (
	PaymentUnpaid   PaymentStatus = "unpaid"
	PaymentPaid     PaymentStatus = "paid"
	PaymentFailed   PaymentStatus = "failed"
	PaymentRefunded PaymentStatus = "refunded"
)
