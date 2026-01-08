package dto

type PaymentGatewayRequest struct {
	OrderID     string
	GrossAmount int64
	Name        string
	Email       string
	PhoneNumber string
}

type PaymentGatewayResponse struct {
	RedirectURL string `json:"redirect_url"`
	Token       string `json:"token"`
}

type PaymentGatewayCallbackRequest struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status" binding:"required"`
	FraudStatus       string `json:"fraud_status" binding:"required"`
}
