package middleware

import (
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// CSRF protects authenticated browser mutations by requiring an explicit
// same-origin Origin or Referer header. Browsers attach these headers to
// state-changing requests, while cross-site forms and scripts cannot forge a
// matching value.
func CSRF(c *fiber.Ctx) error {
	method := c.Method()
	if method == fiber.MethodGet || method == fiber.MethodHead || method == fiber.MethodOptions {
		return c.Next()
	}

	requestOrigin := requestOrigin(c)
	if requestOrigin == "" || requestOrigin != expectedOrigin(c) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "CSRF validation failed",
		})
	}

	return c.Next()
}

func requestOrigin(c *fiber.Ctx) string {
	if origin := strings.TrimSpace(c.Get("Origin")); origin != "" {
		return normalizeOrigin(origin)
	}
	if referer := strings.TrimSpace(c.Get("Referer")); referer != "" {
		return normalizeOrigin(referer)
	}
	return ""
}

func expectedOrigin(c *fiber.Ctx) string {
	scheme := strings.TrimSpace(c.Get("X-Forwarded-Proto"))
	if comma := strings.IndexByte(scheme, ','); comma >= 0 {
		scheme = strings.TrimSpace(scheme[:comma])
	}
	if scheme == "" {
		scheme = "http"
		if c.Protocol() == "https" {
			scheme = "https"
		}
	}
	return scheme + "://" + c.Hostname()
}

func normalizeOrigin(value string) string {
	u, err := url.Parse(value)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ""
	}
	return strings.ToLower(u.Scheme) + "://" + strings.ToLower(u.Host)
}
