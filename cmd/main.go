package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hamwiwatsapon/train-booking-go/internal/application/services"
	"github.com/hamwiwatsapon/train-booking-go/internal/infrastructure/database"
	"github.com/hamwiwatsapon/train-booking-go/internal/infrastructure/repository"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/handlers"
	"github.com/hamwiwatsapon/train-booking-go/internal/presentation/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS for localhost:3000
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
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

	// Initialize train repository, service, and handler
	trainRepo := repository.NewTrainRepository(db)
	trainService := services.NewTrainService(trainRepo)
	trainHandler := handlers.NewTrainHandler(trainService)

	// Add enhanced logging middleware
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)
		log.Printf("[%s] %s %s %d %s", c.IP(), c.Method(), c.Path(), c.Response().StatusCode(), duration)
		return err
	})

	v1 := app.Group("/api/v1")

	// Setup routes
	routes.SetupAuthRoutes(v1, authHandler)
	routes.SetupProfileRoutes(v1, authHandler)
	routes.SetupStationRoutes(v1, trainHandler)

	// Start the server
	app.Listen(":4444")
}
