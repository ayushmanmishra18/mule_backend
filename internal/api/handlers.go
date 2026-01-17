package api

import (
	"log"
	"muleshield/internal/ml"
	"muleshield/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func CheckRisk(c *fiber.Ctx) error {
	var tx models.Transaction
	if err := c.BodyParser(&tx); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid transaction data"})
	}

	// Requesting real-time risk score from ML model
	risk, err := ml.GetRiskScore(tx)
	if err != nil {
		log.Printf("ML Service Error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Fraud detection engine unavailable"})
	}

	// For the prototype, we log suspicious activity to the console instead of a database
	if risk.RiskScore > 0.8 {
		log.Printf("[AUDIT] Flagged Transaction ID: %s, Score: %f", tx.TransactionID, risk.RiskScore)
	}

	return c.JSON(risk)
}

