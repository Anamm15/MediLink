package dto

type PaymentGatewayRequest struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}

type PaymentGatewayResponse struct {
	RedirectURL string `json:"redirect_url"`
}
