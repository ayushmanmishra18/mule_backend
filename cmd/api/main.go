func main() {
    app := fiber.New()

    // Public Route
    app.Post("/login", handlers.Login)

    // Protected Routes
    api := app.Group("/api", api.AuthMiddleware)
    api.Post("/check-risk", handlers.HandleTransaction)

    app.Listen(":3000")
}