package main

import (
	"log"
	"os"
	"muleshield/internal/api"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	// RESTful API Endpoint for financial platforms
	app.Post("/api/check-risk", api.CheckRisk)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))

}




