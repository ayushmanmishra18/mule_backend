package main

import (
	"log"
	"os"
	"muleshield/internal/api"
	"muleshield/internal/audit"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	// Initialize Audit Logger
	auditLogger := audit.NewLogger(os.Getenv("AWS_REGION"), os.Getenv("DYNAMODB_TABLE"))
	h := &api.Handler{Audit: auditLogger}

	// Routes
	app.Post("/login", h.Login)
	
	protected := app.Group("/api", api.AuthMiddleware)
	protected.Post("/check-risk", h.CheckRisk)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}