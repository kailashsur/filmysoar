package admin

import (
	"filmyfly-go-fiber/internal/cache"
	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetSettings displays settings page
func GetSettings(c *fiber.Ctx) error {
	success := c.Query("success", "")
	errorMsg := c.Query("error", "")

	// Fetch all settings
	var settingsData []models.Setting
	result := database.DB.Find(&settingsData)

	if result.Error != nil {
		return c.Render("admin/settings", fiber.Map{
			"title":    "Site Settings",
			"settings": make(map[string]string),
			"success":  "",
			"error":    "Failed to fetch settings: " + result.Error.Error(),
			"user": map[string]interface{}{
				"email": "admin@filmyfly.work",
			},
		})
	}

	// Convert to map for easier template access
	settings := make(map[string]string)
	for _, s := range settingsData {
		settings[s.Key] = s.Value
	}

	return c.Render("admin/settings", fiber.Map{
		"title":    "Site Settings",
		"settings": settings,
		"success":  success,
		"error":    errorMsg,
		"user": map[string]interface{}{
			"email": "admin@filmyfly.work",
		},
	})
}

// PostSettings handles settings update
func PostSettings(c *fiber.Ctx) error {
	// Get all form values
	formData := make(map[string]string)
	c.Request().PostArgs().VisitAll(func(key, value []byte) {
		formData[string(key)] = string(value)
	})

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		for key, value := range formData {
			setting := models.Setting{Key: key, Value: value}
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "key"}},
				DoUpdates: clause.AssignmentColumns([]string{"value", "updatedAt"}),
			}).Create(&setting).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return c.Redirect("/admin/settings?error=Failed to update settings")
	}
	cache.Delete("api:categories")
	cache.DeletePrefix("api:home:")

	return c.Redirect("/admin/settings?success=Settings updated successfully")
}
