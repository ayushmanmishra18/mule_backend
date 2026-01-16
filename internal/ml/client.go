package ml

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"muleshield/pkg/models"
)

func GetRiskScore(tx models.Transaction) (*models.RiskResponse, error) {
	jsonData, _ := json.Marshal(tx)
	resp, err := http.Post(os.Getenv("ML_SERVICE_URL"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.RiskResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return &result, nil
}