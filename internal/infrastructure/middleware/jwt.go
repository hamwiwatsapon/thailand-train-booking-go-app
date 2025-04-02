package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Extract claims and set them in the context (optional)
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Locals("userID", claims["user_id"])
		c.Locals("role", claims["role"])
	}

	// Proceed to the next handler
	return c.Next()
}
