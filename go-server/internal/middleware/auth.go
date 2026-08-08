package middleware

import (
	"filmyfly-go-fiber/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

// InitSession initializes the session store
func InitSession(cfg *config.Config) {
	Store = session.New(session.Config{
		KeyLookup:      "cookie:session_id",
		CookieSecure:   cfg.Environment == "production",
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	})
}

// GetSession returns the session for the current request
func GetSession(c *fiber.Ctx) *session.Session {
	sess, _ := Store.Get(c)
	return sess
}

// RequireAuth middleware checks if user is authenticated
func RequireAuth(c *fiber.Ctx) error {
	sess := GetSession(c)

	// Check if user is authenticated
	uid := sess.Get("uid")
	if uid == nil {
		// Redirect to login for HTML requests
		if c.Get("Accept") == "" || c.Get("Accept") == "text/html" {
			return c.Redirect("/admin/login?error=Please login to continue")
		}
		// Return JSON error for API requests
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized - Please login",
		})
	}

	// Set user info in locals for handlers to access
	c.Locals("user", fiber.Map{
		"uid":         sess.Get("uid"),
		"email":       sess.Get("email"),
		"displayName": sess.Get("displayName"),
	})

	return c.Next()
}

// RedirectIfAuthenticated redirects to dashboard if already logged in
func RedirectIfAuthenticated(c *fiber.Ctx) error {
	sess := GetSession(c)

	uid := sess.Get("uid")
	if uid != nil {
		return c.Redirect("/admin")
	}

	return c.Next()
}
