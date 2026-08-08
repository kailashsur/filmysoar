package routes

import (
	"filmyfly-go-fiber/internal/handlers/admin"
	"filmyfly-go-fiber/internal/handlers/api"
	"filmyfly-go-fiber/internal/handlers/public"
	"filmyfly-go-fiber/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// Setup configures all application routes
func Setup(app *fiber.App) {
	// API Routes (for Astro frontend)
	apiGroup := app.Group("/api")
	{
		// Homepage data
		apiGroup.Get("/home", api.GetHomePageData)

		// Movies
		apiGroup.Get("/movies", api.GetMovies)
		apiGroup.Get("/movies/trending", api.GetTrendingMovies)
		apiGroup.Get("/movies/:slug", api.GetMovieBySlug)

		// Categories
		apiGroup.Get("/categories", api.GetCategories)
		apiGroup.Get("/categories/:slug", api.GetCategoryBySlug)

		// Search
		apiGroup.Get("/search", api.SearchMovies)

		// Static pages
		apiGroup.Get("/static-pages/:slug", api.GetStaticPageBySlug)

		// Astro Settings
		apiGroup.Get("/astro-settings", api.GetAstroSettings)
	}

	// Admin Routes
	adminGroup := app.Group("/admin")
	{
		// Public routes (no authentication required)
		adminGroup.Get("/login", middleware.RedirectIfAuthenticated, admin.GetAdminLogin)
		adminGroup.Post("/login", admin.PostAdminLogin)
		adminGroup.Post("/logout", admin.PostAdminLogout)

		// Protected routes (authentication required)
		adminGroup.Use(middleware.RequireAuth)
		adminGroup.Use(middleware.CSRF)

		adminGroup.Get("/", admin.GetAdminDashboard)
		adminGroup.Get("/system-check", admin.GetSystemCheck)

		// Settings
		adminGroup.Get("/settings", admin.GetSettings)
		adminGroup.Post("/settings", admin.PostSettings)

		// Astro Settings
		adminGroup.Get("/astro-settings", admin.GetAstroSettings)
		adminGroup.Post("/astro-settings", admin.PostAstroSettings)

		// Movie Management
		adminGroup.Get("/movies", admin.GetMovieList)
		adminGroup.Get("/movies/add", admin.GetAddMovie)
		adminGroup.Post("/movies/add", admin.PostAddMovie)
		adminGroup.Get("/movies/bulk-add", admin.GetBulkAddMovies)
		adminGroup.Post("/movies/bulk-add", admin.PostBulkAddMovies)
		adminGroup.Get("/movies/edit/:id", admin.GetEditMovie)
		adminGroup.Post("/movies/edit/:id", admin.PostEditMovie)
		adminGroup.Post("/movies/delete/:id", admin.DeleteMovie)

		// Trending Movies
		adminGroup.Post("/movies/trending/add/:id", admin.AddToTrending)
		adminGroup.Post("/movies/trending/remove/:id", admin.RemoveFromTrending)

		// Static Pages Management
		adminGroup.Get("/static-pages", admin.GetStaticPageList)
		adminGroup.Get("/static-pages/add", admin.GetAddStaticPage)
		adminGroup.Post("/static-pages/add", admin.PostAddStaticPage)
		adminGroup.Get("/static-pages/edit/:id", admin.GetEditStaticPage)
		adminGroup.Post("/static-pages/edit/:id", admin.PostEditStaticPage)
		adminGroup.Post("/static-pages/delete/:id", admin.DeleteStaticPage)

		// Logs Management
		adminGroup.Get("/logs", admin.GetLogs)
		adminGroup.Get("/logs/data", admin.GetLogsData)
		adminGroup.Post("/logs/clear", admin.ClearLogs)
		adminGroup.Get("/logs/download", admin.DownloadLogs)
	}

	// Public Routes
	app.Get("/", public.GetHomePage)
	app.Get("/docs", public.GetDocs)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "Route not found",
		})
	})
}
