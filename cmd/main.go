package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamwiwatsapon/train-booking-go/internal/application/services"
	"github.com/hamwiwatsapon/train-booking-go/internal/infrastructure/repository"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/handlers"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Initialize repository, service, and handler
	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Setup routes
	routes.SetupRoutes(app, authHandler)

	// Start the server
	app.Listen(":3000")
}
