package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/latoulicious/GoShop/config"
	"github.com/latoulicious/GoShop/database"
	"github.com/latoulicious/GoShop/utils"
)

// Declare dsn at the package level
var DSN string

// main is the entry point for the application. It loads configuration, initializes dependencies like the database, sets up routes and middleware, and starts the HTTP server.
func main() {
	// Load configuration from .env file
	config.LoadConfig()

	if os.Getenv("ENVIRONMENT") == "development" {

		log.Println("Checking if JWT secret is set")

		// If JWT_SECRET is not set, generate a new JWT secret key
		if config.JwtSecret == "" {
			log.Println("JWT secret is empty, generating a new one")
			// Generate the JWT secret key
			secretKey, err := utils.GenerateJWTSecretKey()
			if err != nil {
				log.Fatal("Error generating JWT secret key:", err)
			}

			// Print or log the generated JWT secret key
			log.Println("Generated JWT Secret Key:", secretKey)

			// Set the generated JWT secret key in the configuration
			config.JwtSecret = secretKey
		}

	}

	// Initialize Fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Setup CORS
	config.SetupCors(app)

	// Logging Request ID
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	// Connect to the database
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Setup routes
	// router.SetupRoutes(app)

	// Start the server and handle errors
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}

	// Graceful shutdown
	err = app.Shutdown()
	if err != nil {
		log.Println("Graceful shutdown failed:", err)
	}
}
