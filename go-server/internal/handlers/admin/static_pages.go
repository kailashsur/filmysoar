package admin

import (
	"filmyfly-go-fiber/internal/cache"
	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/database/models"

	"github.com/gofiber/fiber/v2"
)

// GetStaticPageList displays list of static pages
func GetStaticPageList(c *fiber.Ctx) error {
	success := c.Query("success", "")
	errorMsg := c.Query("error", "")

	var pages []models.StaticPage
	database.DB.Order("\"createdAt\" DESC").Find(&pages)

	return c.Render("admin/static-pages/list", fiber.Map{
		"title":   "Manage Static Pages",
		"pages":   pages,
		"success": success,
		"error":   errorMsg,
		"user": map[string]interface{}{
			"email": "admin@filmyfly.work",
		},
	})
}

// GetAddStaticPage renders add static page form
func GetAddStaticPage(c *fiber.Ctx) error {
	return c.Render("admin/static-pages/add", fiber.Map{
		"title": "Add Static Page",
		"user":  getUserFromSession(c),
	})
}

// PostAddStaticPage handles static page creation
func PostAddStaticPage(c *fiber.Ctx) error {
	page := new(models.StaticPage)

	if err := c.BodyParser(page); err != nil {
		return c.Redirect("/admin/static-pages/add?error=Invalid data")
	}

	if page.Title == "" || page.Slug == "" || page.Content == "" {
		return c.Redirect("/admin/static-pages/add?error=Title, slug, and content are required")
	}

	// Check if slug already exists
	var existing models.StaticPage
	if err := database.DB.Where("slug = ?", page.Slug).First(&existing).Error; err == nil {
		return c.Redirect("/admin/static-pages/add?error=A page with this slug already exists")
	}

	// Handle isPublished checkbox
	isPublished := c.FormValue("isPublished")
	page.IsPublished = isPublished == "on" || isPublished == "true"

	if err := database.DB.Create(page).Error; err != nil {
		return c.Redirect("/admin/static-pages/add?error=Failed to create page")
	}
	cache.DeletePrefix("api:static-page:")

	return c.Redirect("/admin/static-pages?success=Page created successfully")
}

// GetEditStaticPage renders edit static page form
func GetEditStaticPage(c *fiber.Ctx) error {
	id := c.Params("id")

	var page models.StaticPage
	if err := database.DB.First(&page, id).Error; err != nil {
		return c.Redirect("/admin/static-pages?error=Page not found")
	}

	return c.Render("admin/static-pages/edit", fiber.Map{
		"title": "Edit Static Page",
		"page":  page,
		"user":  getUserFromSession(c),
	})
}

// PostEditStaticPage handles static page update
func PostEditStaticPage(c *fiber.Ctx) error {
	id := c.Params("id")

	var page models.StaticPage
	if err := database.DB.First(&page, id).Error; err != nil {
		return c.Redirect("/admin/static-pages?error=Page not found")
	}

	if err := c.BodyParser(&page); err != nil {
		return c.Redirect("/admin/static-pages/edit/" + id + "?error=Invalid data")
	}

	if page.Title == "" || page.Slug == "" || page.Content == "" {
		return c.Redirect("/admin/static-pages/edit/" + id + "?error=Title, slug, and content are required")
	}

	// Check if slug is taken by another page
	var existing models.StaticPage
	if err := database.DB.Where("slug = ? AND id != ?", page.Slug, page.ID).First(&existing).Error; err == nil {
		return c.Redirect("/admin/static-pages/edit/" + id + "?error=A page with this slug already exists")
	}

	// Handle isPublished checkbox
	isPublished := c.FormValue("isPublished")
	page.IsPublished = isPublished == "on" || isPublished == "true"

	if err := database.DB.Save(&page).Error; err != nil {
		return c.Redirect("/admin/static-pages/edit/" + id + "?error=Failed to update page")
	}
	cache.DeletePrefix("api:static-page:")

	return c.Redirect("/admin/static-pages?success=Page updated successfully")
}

// DeleteStaticPage handles static page deletion
func DeleteStaticPage(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := database.DB.Delete(&models.StaticPage{}, id).Error; err != nil {
		return c.Redirect("/admin/static-pages?error=Failed to delete page")
	}
	cache.DeletePrefix("api:static-page:")

	return c.Redirect("/admin/static-pages?success=Page deleted successfully")
}
