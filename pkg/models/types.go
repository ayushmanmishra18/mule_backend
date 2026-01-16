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
	RiskScore      float64 `json:"risk_score"`
	Recommendation string  `json:"recommendation"`
}

type Credentials struct {
	BankID   string `json:"bank_id"`
	Password string `json:"password"`
}