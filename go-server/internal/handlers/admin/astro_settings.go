package admin

import (
	"filmyfly-go-fiber/internal/cache"
	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetAstroSettings displays Astro settings page
func GetAstroSettings(c *fiber.Ctx) error {
	success := c.Query("success", "")
	errorMsg := c.Query("error", "")

	// Fetch all Astro settings
	var astroSettingsData []models.AstroSetting
	if err := database.DB.Find(&astroSettingsData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("admin/astro-settings", fiber.Map{
			"title":    "Astro Website Settings",
			"settings": make(map[string]string),
			"success":  "",
			"error":    "Failed to fetch Astro settings",
			"user":     getUserFromSession(c),
		})
	}

	// Convert to map for easier template access
	settings := make(map[string]string)
	for _, s := range astroSettingsData {
		settings[s.Key] = s.Value
	}

	return c.Render("admin/astro-settings", fiber.Map{
		"title":    "Astro Website Settings",
		"settings": settings,
		"success":  success,
		"error":    errorMsg,
		"user":     getUserFromSession(c),
	})
}

// PostAstroSettings handles Astro settings update
func PostAstroSettings(c *fiber.Ctx) error {
	// Get all form values
	formData := make(map[string]string)
	c.Request().PostArgs().VisitAll(func(key, value []byte) {
		formData[string(key)] = string(value)
	})

	// Upsert all settings in one transaction, so the page never reports a
	// successful update when one of its values could not be stored.
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		for key, value := range formData {
			setting := models.AstroSetting{Key: key, Value: value}
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "key"}},
				DoUpdates: clause.AssignmentColumns([]string{"value", "updatedAt"}),
			}).Create(&setting).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return c.Redirect("/admin/astro-settings?error=Failed to update Astro settings")
	}
	cache.Delete("api:astro-settings")

	return c.Redirect("/admin/astro-settings?success=Astro settings updated successfully")
}
