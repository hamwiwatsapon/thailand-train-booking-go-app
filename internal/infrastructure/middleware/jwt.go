package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Secret key for signing JWT tokens
var secretKey = []byte("!%@2ferGE525s2544423BV23@_(sfECe@4221Seipklx.nhHIpeo)") // Replace with your actual secret key

// JWTMiddleware configures the JWT middleware
func JWTMiddleware(c *fiber.Ctx) error {
	tokenHeader := c.Get("Authorization")
	tokenString := tokenHeader[len("Bearer "):]
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid token",
		})
	}
	return c.Next()
}
