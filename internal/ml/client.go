
func GetRiskScore(tx models.Transaction) (*models.RiskResponse, error) {
    // Calls your Python/FastAPI service
    resp, err := http.Post(os.Getenv("ML_URL")+"/predict", "application/json", body)
    // ... decode response into RiskResponse
}