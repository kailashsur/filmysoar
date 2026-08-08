package main

import (
	"fmt"
	"log"

	"filmyfly-go-fiber/internal/config"
	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/middleware"
	"filmyfly-go-fiber/internal/routes"
	"filmyfly-go-fiber/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize logger
	utils.InitLogger()
	utils.Info("Starting FilmyFly Go Fiber server...")

	// Load configuration
	cfg := config.Load()
	utils.Info("Configuration loaded")

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	utils.Info("Database connected successfully")

	// Initialize Firebase
	if err := config.InitializeFirebase(); err != nil {
		utils.Warn("Firebase initialization failed: %v. Admin authentication will not work.", err)
	} else {
		utils.Info("Firebase initialized successfully")
	}

	// Initialize session store
	middleware.InitSession(cfg)
	utils.Info("Session store initialized")

	// Initialize template engine
	engine := html.New("./views", ".html")
	engine.Reload(cfg.Environment == "development")
	engine.Debug(cfg.Environment == "development")

	// Add template functions
	engine.AddFunc("add", func(a, b int) int {
		return a + b
	})
	engine.AddFunc("sub", func(a, b int) int {
		return a - b
	})

	// Add pipeline versions for use with |
	engine.AddFuncMap(map[string]interface{}{
		"add": func(b int, a int) int { return a + b },
		"sub": func(b int, a int) int { return a - b },
	})

	// Create Fiber app
	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: middleware.ErrorHandler,
		ServerHeader: "FilmyFly",
		AppName:      "FilmyFly v1.0",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))
	app.Use(middleware.CORS(cfg))

	// Static files
	app.Static("/", "./public")

	// Settings middleware (load settings for all views)
	app.Use(func(c *fiber.Ctx) error {
		var settings []struct {
			Key   string
			Value string
		}

		if err := database.DB.Table("settings").Select("\"key\", value").Scan(&settings).Error; err != nil {
			return fmt.Errorf("load site settings: %w", err)
		}

		settingsObj := make(map[string]string)
		for _, s := range settings {
			settingsObj[s.Key] = s.Value
		}

		c.Locals("settings", settingsObj)
		return c.Next()
	})

	// Setup routes
	routes.Setup(app)

	// Start server
	port := cfg.Port
	utils.Info("Server starting on port %s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
