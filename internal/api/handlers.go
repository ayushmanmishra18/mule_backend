func HandleTransaction(c *fiber.Ctx) error {
    var tx models.Transaction
    if err := c.BodyParser(&tx); err != nil {
        return c.Status(400).SendString("Invalid Input")
    }

    risk, _ := ml.GetRiskScore(tx) // Call the ML service
    
    // Optional: Log high-risk transactions to DynamoDB for audit
    if risk.RiskScore > 0.7 {
        audit.LogToDynamoDB(tx, risk)
    }

    return c.JSON(risk)
}