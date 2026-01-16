package api

import (
    "os"
    "strings"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if token != nil && token.Valid {
        return c.Next()
    }
    return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
}