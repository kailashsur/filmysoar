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

// GetCategories handles GET /api/categories
func GetCategories(c *fiber.Ctx) error {
	if value, ok := cache.Get("api:categories"); ok {
		return c.JSON(value)
	}
	// Get categories with movie counts
	type CategoryWithCount struct {
		models.Category
		MovieCount int64 `json:"movieCount"`
	}

	var results []CategoryWithCount

	if err := database.DB.
		Model(&models.Category{}).
		Select("categories.*, COUNT(movies.id) as movie_count").
		Joins("LEFT JOIN movies ON movies.\"categoryId\" = categories.id").
		Group("categories.id").
		Order("categories.name ASC").
		Scan(&results).Error; err != nil {
		utils.Error("Failed to fetch categories: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch categories",
		})
	}

	// Transform to match Node.js API response format
	response := make([]fiber.Map, len(results))
	for i, cat := range results {
		response[i] = fiber.Map{
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

	responseBody := fiber.Map{
		"success": true,
		"data":    response,
	}
	cache.Set("api:categories", responseBody, time.Minute)
	return c.JSON(responseBody)
}

// GetCategoryBySlug handles GET /api/categories/:slug
func GetCategoryBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", strconv.Itoa(ItemsPerPage)))
	page, limit = utils.NormalizePagination(page, limit, ItemsPerPage, MaxItemsPerPage)
	offset := utils.GetOffset(page, limit)

	// Find category by slug
	var category models.Category
	if err := database.DB.Where("slug = ?", slug).First(&category).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "Category not found",
		})
	}

	// Get movies in this category
	var movies []models.Movie
	var total int64

	database.DB.Model(&models.Movie{}).Where("\"categoryId\" = ?", category.ID).Count(&total)

	if err := database.DB.
		Select("id, title, slug, thumbnail, keywords, \"releaseYear\", genre").
		Where("\"categoryId\" = ?", category.ID).
		Order("\"createdAt\" DESC").
		Offset(offset).
		Limit(limit).
		Find(&movies).Error; err != nil {
		utils.Error("Failed to fetch movies: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch category movies",
		})
	}

	pagination := utils.NewPagination(page, limit, total)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"category":   category,
			"movies":     movies,
			"pagination": pagination,
		},
	})
}
