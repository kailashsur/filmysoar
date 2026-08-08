package api

import (
	"strconv"
	"time"

	"filmyfly-go-fiber/internal/cache"
	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/database/models"
	"filmyfly-go-fiber/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// GetHomePageData handles GET /api/home
func GetHomePageData(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	page = utils.NormalizePage(page)
	limit := 50 // Match the original 50 movies per page
	offset := utils.GetOffset(page, limit)
	cacheKey := "api:home:" + strconv.Itoa(page)
	if value, ok := cache.Get(cacheKey); ok {
		return c.JSON(value)
	}

	// Fetch trending movies
	var trendingMoviesData []models.TrendingMovie
	database.DB.
		Preload("Movie").
		Order("\"order\" ASC").
		Find(&trendingMoviesData)

	trendingMovies := make([]models.Movie, 0, len(trendingMoviesData))
	for _, tm := range trendingMoviesData {
		trendingMovies = append(trendingMovies, tm.Movie)
	}

	// Fetch recent movies
	var recentMovies []models.Movie
	if err := database.DB.
		Select("id, title, slug, thumbnail, keywords, \"releaseYear\", genre").
		Order("\"createdAt\" DESC").
		Offset(offset).
		Limit(limit).
		Find(&recentMovies).Error; err != nil {
		utils.Error("Failed to fetch recent movies: %v", err)
		recentMovies = []models.Movie{} // Return empty array instead of nil
	}

	// Initialize as empty array if nil
	if recentMovies == nil {
		recentMovies = []models.Movie{}
	}

	// Fetch categories with counts
	type CategoryWithCount struct {
		models.Category
		MovieCount int64 `json:"movieCount"`
	}

	var categoryResults []CategoryWithCount
	database.DB.
		Model(&models.Category{}).
		Select("categories.*, COUNT(movies.id) as movie_count").
		Joins("LEFT JOIN movies ON movies.\"categoryId\" = categories.id").
		Group("categories.id").
		Order("categories.name ASC").
		Scan(&categoryResults)

	// Transform categories to match Node.js format
	categories := make([]fiber.Map, len(categoryResults))
	for i, cat := range categoryResults {
		categories[i] = fiber.Map{
			"id":          cat.ID,
			"name":        cat.Name,
			"slug":        cat.Slug,
			"description": cat.Description,
			"createdAt":   cat.CreatedAt,
			"updatedAt":   cat.UpdatedAt,
			"_count": fiber.Map{
				"movies": cat.MovieCount,
			},
		}
	}

	// Get total count for pagination
	var total int64
	database.DB.Model(&models.Movie{}).Count(&total)

	pagination := utils.NewPagination(page, limit, total)

	response := fiber.Map{
		"success": true,
		"data": fiber.Map{
			"trendingMovies": trendingMovies,
			"recentMovies":   recentMovies,
			"categories":     categories,
			"pagination":     pagination,
		},
	}
	cache.Set(cacheKey, response, 30*time.Second)
	return c.JSON(response)
}

// GetStaticPageBySlug handles GET /api/static-pages/:slug
func GetStaticPageBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	cacheKey := "api:static-page:" + slug
	if value, ok := cache.Get(cacheKey); ok {
		return c.JSON(value)
	}

	// Try with the slug as-is first
	var page models.StaticPage
	err := database.DB.
		Where("slug = ? AND \"isPublished\" = ?", slug, true).
		First(&page).Error

	// If not found, try with .html suffix
	if err != nil {
		slugWithHTML := slug + ".html"
		err = database.DB.
			Where("slug = ? AND \"isPublished\" = ?", slugWithHTML, true).
			First(&page).Error
	}

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "Page not found",
		})
	}

	response := fiber.Map{
		"success": true,
		"data":    page,
	}
	cache.Set(cacheKey, response, time.Minute)
	return c.JSON(response)
}

// GetAstroSettings handles GET /api/astro-settings
func GetAstroSettings(c *fiber.Ctx) error {
	if value, ok := cache.Get("api:astro-settings"); ok {
		return c.JSON(value)
	}
	var settings []models.AstroSetting

	if err := database.DB.
		Select("\"key\", value").
		Find(&settings).Error; err != nil {
		utils.Error("Failed to fetch Astro settings: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch Astro settings",
		})
	}

	// Convert to key-value object
	settingsObj := make(map[string]string)
	for _, setting := range settings {
		settingsObj[setting.Key] = setting.Value
	}

	response := fiber.Map{
		"success": true,
		"data":    settingsObj,
	}
	cache.Set("api:astro-settings", response, time.Minute)
	return c.JSON(response)
}
