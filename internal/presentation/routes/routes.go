package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamwiwatsapon/train-booking-go/internal/infrastructure/middleware"
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
	profile := app.Group("/auth")

	profile.Get("/profile", func(c *fiber.Ctx) error {
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

func SetupStationRoutes(app fiber.Router, trainHandler *handlers.TrainHandler) {

	// Public station routes (no middleware)
	station := app.Group("/stations")
	station.Get("/", trainHandler.GetStations)

	// Protected station routes (with middleware)
	auth := app.Group("/auth/stations", middleware.JWTMiddleware)
	auth.Post("/bulk", trainHandler.BulkCreateStation)

	// Protected station type routes
	station.Get("/type", trainHandler.GetStationTypes)
	auth.Post("/type", trainHandler.CreateStationType)
	auth.Put("/type", trainHandler.UpdateStationType)
	auth.Delete("/type", trainHandler.DeleteStationType)
}
