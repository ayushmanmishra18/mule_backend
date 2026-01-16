package models

type Transaction struct {
    TransactionID string  `json:"transaction_id"`
    SenderAccount string  `json:"sender_account"`
    ReceiverID    string  `json:"receiver_id"`
    Amount        float64 `json:"amount"`
    Timestamp     string  `json:"timestamp"`
}

type RiskResponse struct {
    RiskScore      float64 `json:"risk_score"`
    Recommendation string  `json:"recommendation"` // ALLOW, FLAG, BLOCK
}

type BankUser struct {
    BankID   string `json:"bank_id"`
    Password string `json:"password"`
}