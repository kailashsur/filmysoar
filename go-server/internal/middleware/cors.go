package middleware

import (
	"filmyfly-go-fiber/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CORS returns CORS middleware configured for the application
func CORS(cfg *config.Config) fiber.Handler {
	allowedOrigins := []string{
		
		"https://filmyfly.work",
		"https://www.filmyfly.work",
		"https://filmyflyhd.space",
		"https://www.filmyflyhd.space",
		"https://filmyflyhd.online",
		"https://filmysoar.online",
		"https://www.filmysoar.online",
		
	}

	if cfg.FrontendURL != "" {
		allowedOrigins = append(allowedOrigins, cfg.FrontendURL)
	}

	return cors.New(cors.Config{
		AllowOrigins:     joinStrings(allowedOrigins, ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization",
	})
}

func joinStrings(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
