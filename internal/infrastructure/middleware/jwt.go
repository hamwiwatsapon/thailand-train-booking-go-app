package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hamwiwatsapon/train-booking-go/internal/application/services"
)

// Secret key for signing JWT tokens
var secretKey = []byte("!%@2ferGE525s2544423BV23@_(sfECe@4221Seipklx.nhHIpeo)") // Replace with your actual secret key

// JWTMiddleware validates the JWT token from the Authorization header
func JWTMiddleware(c *fiber.Ctx) error {
	// Get the Authorization header
	tokenHeader := c.Get("Authorization")
	if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid Authorization header",
		})
	}

	// Extract the token string
	tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
	// Parse and validate the token
	token, err := services.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user ID and role from the claims
		userIDFloat, ok := claims["user"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid user ID in token",
			})
		}

		// Log the extracted user ID for debugging
		fmt.Printf("Extracted user ID from token: %v\n", userIDFloat)

		userID := uint(userIDFloat)
		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid role in token",
			})
		}

		// Log the role for debugging
		fmt.Printf("Extracted role from token: %v\n", role)

		// Store user ID, email, and role in context locals for later use
		c.Locals("user", userID)
		c.Locals("email", claims["email"])
		c.Locals("role", role)
	}

	// Proceed to the next handler
	return c.Next()
}
