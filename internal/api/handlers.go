package api

import (
	"fmt"
	"log"
	"muleshield/internal/blockchain"
	"muleshield/internal/ml"
	"muleshield/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// Handler struct ab Blockchain service ko hold karega
type Handler struct {
	Blockchain *blockchain.Service
}

// Constructor
func NewHandler(bc *blockchain.Service) *Handler {
	return &Handler{Blockchain: bc}
}

// CheckRisk ab Handler ka method ban gaya hai
func (h *Handler) CheckRisk(c *fiber.Ctx) error {
	var tx models.Transaction
	if err := c.BodyParser(&tx); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid transaction data"})
	}

	// --- NEW: LAYER 1 - BLOCKCHAIN CHECK (The "Hard" Block) ---
	log.Println(" Checking Blockchain for Blacklisted Accounts...")
	
	senderIsMule, _ := h.Blockchain.IsWalletBlacklisted(tx.SenderAccount)
	receiverIsMule, _ := h.Blockchain.IsWalletBlacklisted(tx.ReceiverID)

	if senderIsMule || receiverIsMule {
		log.Printf(" [BLOCKED] Known Mule Detected! Sender: %v, Receiver: %v", senderIsMule, receiverIsMule)
		
		// Return immediately - DO NOT call ML
		return c.JSON(models.RiskResponse{
			TransactionID:  tx.TransactionID,
			RiskScore:      1.0, // Maximum Risk
			Recommendation: "BLOCK",
			//Reason:         "Account found in Blockchain Blacklist",
		})
	}

	// --- LAYER 2 - ML CHECK (The "Soft" Block) ---
	// (This only runs if the user was NOT found on blockchain)
	risk, err := ml.GetRiskScore(tx)
	if err != nil {
		log.Printf("ML Service Error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Fraud detection engine unavailable"})
	}

	// If ML says High Risk -> Report to Blockchain
	if risk.RiskScore > 0.8 {
		log.Printf("🚨 [AUDIT] HIGH RISK DETECTED (%f). Reporting to Blockchain...", risk.RiskScore)
		go func() {
			h.Blockchain.ReportMule(
				tx.TransactionID,
				tx.SenderAccount,
				tx.ReceiverID,
				fmt.Sprintf("%.0f", tx.Amount),
				tx.Location,
				fmt.Sprintf("High ML Risk Score: %f", risk.RiskScore),
			)
		}()
	}

	return c.JSON(risk)
}