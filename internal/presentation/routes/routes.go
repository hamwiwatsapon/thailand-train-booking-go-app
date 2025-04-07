package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamwiwatsapon/train-booking-go/internal/infrastructure/middleware"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/handlers"
)

func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	v1 := app.Group("/api/v1")

	// Authentication routes
	v1.Post("/register", authHandler.Register)
	v1.Post("/login", authHandler.Login)
	v1.Post("/refresh-token", authHandler.Refresh)
	v1.Post("/check-user", authHandler.CheckUser)
	v1.Post("/otp-login", authHandler.OTPLogin)

	protected := v1.Group("/protected", middleware.JWTMiddleware)
	protected.Get("/profile", func(c *fiber.Ctx) error {
		userID := c.Locals("userID")
		role := c.Locals("role")
		userEmail := c.Locals("userEmail")
		return c.JSON(fiber.Map{
			"user_id": userID,
			"email":   userEmail,
			"role":    role,
		})
	})
}
