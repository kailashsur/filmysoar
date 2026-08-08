package public

import (
	"filmyfly-go-fiber/internal/database"
	"filmyfly-go-fiber/internal/database/models"

	"filmyfly-go-fiber/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// GetHomePage handles GET /
func GetHomePage(c *fiber.Ctx) error {
	// Fetch trending movies
	var trendingMoviesData []models.TrendingMovie
	if err := database.DB.
		Preload("Movie").
		Order("\"order\" ASC").
		Find(&trendingMoviesData).Error; err != nil {
		utils.Error("Failed to fetch trending movies: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}

	trendingMovies := make([]models.Movie, 0, len(trendingMoviesData))
	for _, tm := range trendingMoviesData {
		trendingMovies = append(trendingMovies, tm.Movie)
	}

	// Fetch recent movies
	var recentMovies []models.Movie
	if err := database.DB.
		Select("id, title, slug, thumbnail, keywords, \"releaseYear\", genre").
		Order("\"createdAt\" DESC").
		Limit(50).
		Find(&recentMovies).Error; err != nil {
		utils.Error("Failed to fetch recent movies: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}

	// Fetch categories
	var categories []models.Category
	if err := database.DB.Order("name ASC").Find(&categories).Error; err != nil {
		utils.Error("Failed to fetch categories: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}

	return c.Render("index", fiber.Map{
		"title":          "FilmyFly - Download Latest Movies",
		"trendingMovies": trendingMovies,
		"recentMovies":   recentMovies,
		"categories":     categories,
	})
}
