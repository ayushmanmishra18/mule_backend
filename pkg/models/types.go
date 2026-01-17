package models

type Transaction struct {
	TransactionID string  `json:"transaction_id"`
	SenderAccount string  `json:"sender_account"`
	ReceiverID    string  `json:"receiver_id"`
	Amount        float64 `json:"amount"`
	Timestamp     string  `json:"timestamp"`
	Location      string  `json:"location"`
}

type RiskResponse struct {
	TransactionID  string  `json:"transaction_id"`
	RiskScore      float64 `json:"risk_score"` // 0.0 to 1.0
	Recommendation string  `json:"recommendation"` // ALLOW, FLAG, BLOCK
}