package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupCors configures CORS for the application.
// It creates a CORS middleware instance with an AllowOriginsFunc that checks
// the ENVIRONMENT variable to determine if CORS should allow all origins
// (development) or deny all (production).

func SetupCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))
}
