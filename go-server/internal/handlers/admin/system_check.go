package admin

import (
	"filmyfly-go-fiber/internal/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetSystemCheck displays system status page
func GetSystemCheck(c *fiber.Ctx) error {
	// Get database size
	var size string
	err := database.DB.Raw("SELECT pg_size_pretty(pg_database_size(current_database())) as size").Scan(&size).Error

	if err != nil || size == "" {
		size = "Unknown"
	}

	databaseStatus := map[string]interface{}{
		"size":        size,
		"lastChecked": time.Now().Format("2006-01-02 15:04:05"),
	}

	return c.Render("admin/system-check", fiber.Map{
		"title":          "System Check",
		"databaseStatus": databaseStatus,
		"user":           getUserFromSession(c),
	})
}
