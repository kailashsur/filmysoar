package api

import (
	"strconv"

	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/database/models"
	"filmyfly-go-fiber/internal/utils"

	"github.com/gofiber/fiber/v2"
)

const (
	ItemsPerPage    = 20
	MaxItemsPerPage = 100
)

// GetMovies handles GET /api/movies
func GetMovies(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", strconv.Itoa(ItemsPerPage)))
	page, limit = utils.NormalizePagination(page, limit, ItemsPerPage, MaxItemsPerPage)
	offset := utils.GetOffset(page, limit)

	var movies []models.Movie
	var total int64

	// Get movies with pagination
	if err := database.DB.Model(&models.Movie{}).Count(&total).Error; err != nil {
		utils.Error("Failed to count movies: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch movies",
		})
	}

	if err := database.DB.
		Select("id, title, slug, thumbnail, keywords, \"releaseYear\", genre, \"createdAt\"").
		Order("\"createdAt\" DESC").
		Offset(offset).
		Limit(limit).
		Find(&movies).Error; err != nil {
		utils.Error("Failed to fetch movies: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch movies",
		})
	}

	pagination := utils.NewPagination(page, limit, total)

	return c.JSON(fiber.Map{
		"success":    true,
		"data":       movies,
		"pagination": pagination,
	})
}

// GetTrendingMovies handles GET /api/movies/trending
func GetTrendingMovies(c *fiber.Ctx) error {
	var trendingMovies []models.TrendingMovie

	if err := database.DB.
		Preload("Movie").
		Order("\"order\" ASC").
		Find(&trendingMovies).Error; err != nil {
		utils.Error("Failed to fetch trending movies: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch trending movies",
		})
	}

	// Extract movies from trending
	movies := make([]models.Movie, 0, len(trendingMovies))
	for _, tm := range trendingMovies {
		movies = append(movies, tm.Movie)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    movies,
	})
}

// GetMovieBySlug handles GET /api/movies/:slug
func GetMovieBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var movie models.Movie
	if err := database.DB.
		Preload("Category").
		Preload("Trending").
		Where("slug = ?", slug).
		First(&movie).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "Movie not found",
		})
	}

	// Get related movies (same category or recent)
	var relatedMovies []models.Movie
	query := database.DB.
		Select("id, title, slug, thumbnail, \"releaseYear\", genre").
		Where("id != ?", movie.ID).
		Order("\"createdAt\" DESC").
		Limit(10)

	if movie.CategoryID != nil {
		query = query.Where("\"categoryId\" = ?", *movie.CategoryID)
	}

	query.Find(&relatedMovies)

	// Get download redirect URL from settings
	var downloadRedirectURL string
	var setting models.Setting
	if err := database.DB.Where("key = ?", "downloadRedirectUrl").First(&setting).Error; err == nil {
		downloadRedirectURL = setting.Value
	} else {
		utils.Error("Failed to fetch downloadRedirectUrl setting: %v", err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"movie":               movie,
			"relatedMovies":       relatedMovies,
			"downloadRedirectUrl": downloadRedirectURL,
		},
	})
}
