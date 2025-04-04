package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hamwiwatsapon/train-booking-go/internal/application/services"
	"github.com/hamwiwatsapon/train-booking-go/internal/infrastructure/database"
	"github.com/hamwiwatsapon/train-booking-go/internal/infrastructure/repository"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/handlers"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS for localhost:3000
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Initialize database
	dbInstance := database.NewDatabase()
	db, err := dbInstance.Connect()
	if err != nil {
		dbInstance.Close()
		panic("failed to connect to database")
	}

	dbInstance.Migrate()

	// Initialize repository, service, and handler
	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Add logging middleware
	app.Use(func(c *fiber.Ctx) error {
		println("Request:", c.Method(), c.Path())
		return c.Next()
	})

	// Setup routes
	routes.SetupRoutes(app, authHandler)

	// Start the server
	app.Listen(":4444")
}
