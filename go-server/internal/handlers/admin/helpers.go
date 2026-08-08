package admin

import (
	"filmyfly-go-fiber/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// getUserFromSession is a helper function to get user data from session
func getUserFromSession(c *fiber.Ctx) map[string]interface{} {
	sess := middleware.GetSession(c)

	return map[string]interface{}{
		"email":       sess.Get("email"),
		"displayName": sess.Get("displayName"),
		"uid":         sess.Get("uid"),
	}
}
