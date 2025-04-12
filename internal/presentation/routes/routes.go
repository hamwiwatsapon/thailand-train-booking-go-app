package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/handlers"
)

func SetupAuthRoutes(app fiber.Router, authHandler *handlers.AuthHandler) {
	// Authentication routes
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)
	app.Post("/refresh-token", authHandler.Refresh)
	app.Post("/check-user", authHandler.CheckUser)
	app.Post("/otp-login", authHandler.OTPLogin)
}

func SetupProfileRoutes(app fiber.Router, authHandler *handlers.AuthHandler) {
	// Profile routes
	profile := app.Group("/profile")

	profile.Get("/", func(c *fiber.Ctx) error {
		userID := c.Locals("user")
		userRole := c.Locals("role")
		userEmail := c.Locals("email")

		if userID == nil || userRole == nil || userEmail == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access",
			})
		}

		return c.JSON(fiber.Map{
			"user_id": userID,
			"email":   userEmail,
			"role":    userRole,
		})
	})
}

func SetupTrainRoutes(app fiber.Router, trainHandler *handlers.TrainHandler) {
	// Station routes
	station := app.Group("/stations")
	station.Get("/type", trainHandler.GetTrainTypes)
	station.Post("/type", trainHandler.CreateTrainType)
}
