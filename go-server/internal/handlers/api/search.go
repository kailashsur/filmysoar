package api

import (
	"strconv"
	"strings"

	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/database/models"
	"filmyfly-go-fiber/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// SearchMovies handles GET /api/search
func SearchMovies(c *fiber.Ctx) error {
	query := strings.TrimSpace(c.Query("q", c.Query("to-search", "")))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", strconv.Itoa(ItemsPerPage)))
	page, limit = utils.NormalizePagination(page, limit, ItemsPerPage, MaxItemsPerPage)
	offset := utils.GetOffset(page, limit)

	if query == "" {
		return c.JSON(fiber.Map{
			"success": true,
			"data": fiber.Map{
				"query":      "",
				"movies":     []models.Movie{},
				"pagination": utils.NewPagination(1, limit, 0),
			},
		})
	}

	// Build search condition (case-insensitive, search in title only to avoid NULL issues)
	searchPattern := "%" + strings.ToLower(query) + "%"

	var movies []models.Movie
	var total int64

	// Count total matching movies
	database.DB.Model(&models.Movie{}).
		Where("LOWER(title) LIKE ?", searchPattern).
		Count(&total)

	// Get movies with pagination
	if err := database.DB.
		Where("LOWER(title) LIKE ?", searchPattern).
		Order("\"createdAt\" DESC").
		Offset(offset).
		Limit(limit).
		Find(&movies).Error; err != nil {
		utils.Error("Failed to search movies: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to search movies",
		})
	}

	// Initialize as empty array if nil
	if movies == nil {
		movies = []models.Movie{}
	}

	pagination := utils.NewPagination(page, limit, total)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"query":      query,
			"movies":     movies,
			"pagination": pagination,
		},
	})
}
