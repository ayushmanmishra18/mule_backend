package main

import (
	"log"
	"os"
	"muleshield/internal/api"
	"muleshield/internal/blockchain"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// 1. Initialize Blockchain Service
	bcService, err := blockchain.NewService()
	if err != nil {
		log.Printf(" Warning: Could not connect to Blockchain: %v", err)
	} else {
		log.Println(" Connected to Sepolia Blockchain")
	}

	// 2. Initialize API Handler with dependencies
	handler := api.NewHandler(bcService)

	app := fiber.New()

	// 3. Update Route to use the new handler method
	app.Post("/api/check-risk", handler.CheckRisk)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}