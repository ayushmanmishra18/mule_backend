package api

import (
	"muleshield/internal/auth"
	"muleshield/internal/ml"
	"muleshield/internal/audit"
	"muleshield/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Audit *audit.Logger
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var creds models.Credentials
	c.BodyParser(&creds)

	// Prototype: basic check
	if creds.BankID == "admin_bank" && creds.Password == "secure123" {
		token, _ := auth.GenerateToken(creds.BankID)
		return c.JSON(fiber.Map{"token": token})
	}
	return c.Status(401).SendString("Unauthorized")
}

func (h *Handler) CheckRisk(c *fiber.Ctx) error {
	var tx models.Transaction
	c.BodyParser(&tx)

	risk, err := ml.GetRiskScore(tx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ML Service Unreachable"})
	}

	if risk.RiskScore > 0.8 {
		go h.Audit.LogHighRisk(tx, *risk)
	}

	return c.JSON(risk)
}