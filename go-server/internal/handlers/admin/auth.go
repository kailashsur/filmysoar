package admin

import (
	"context"
	"filmyfly-go-fiber/internal/config"
	"filmyfly-go-fiber/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// GetAdminLogin renders the login page
func GetAdminLogin(c *fiber.Ctx) error {
	cfg := config.Load()

	return c.Render("admin/login", fiber.Map{
		"title": "Admin Login",
		"error": c.Query("error"),
		"firebaseConfig": fiber.Map{
			"apiKey":            cfg.FirebaseAPIKey,
			"authDomain":        cfg.FirebaseAuthDomain,
			"projectId":         cfg.FirebaseProjectID,
			"storageBucket":     cfg.FirebaseStorageBucket,
			"messagingSenderId": cfg.FirebaseMessagingSenderID,
			"appId":             cfg.FirebaseAppID,
		},
	})
}

// PostAdminLogin handles login with Firebase token verification
func PostAdminLogin(c *fiber.Ctx) error {
	var req struct {
		IDToken string `json:"idToken"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request",
		})
	}

	// Verify Firebase ID token
	if config.FirebaseAuth == nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Firebase not initialized",
		})
	}

	token, err := config.FirebaseAuth.VerifyIDToken(context.Background(), req.IDToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid authentication token",
		})
	}

	// Get user info from token claims
	email := ""
	displayName := ""

	if emailClaim, ok := token.Claims["email"].(string); ok {
		email = emailClaim
	}

	emailVerified, _ := token.Claims["email_verified"].(bool)
	if email == "" || !emailVerified {
		return c.Status(403).JSON(fiber.Map{
			"success": false,
			"error":   "A verified email address is required for admin access",
		})
	}

	if config.AppConfig == nil || !config.AppConfig.IsAdminEmail(email) {
		return c.Status(403).JSON(fiber.Map{
			"success": false,
			"error":   "This account is not authorized to access the admin panel",
		})
	}

	if nameClaim, ok := token.Claims["name"].(string); ok {
		displayName = nameClaim
	}

	// If no name in claims, try display_name or use email
	if displayName == "" {
		if displayNameClaim, ok := token.Claims["display_name"].(string); ok {
			displayName = displayNameClaim
		} else {
			displayName = email
		}
	}

	// Create session
	sess := middleware.GetSession(c)
	sess.Set("uid", token.UID)
	sess.Set("email", email)
	sess.Set("displayName", displayName)

	if err := sess.Save(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to create session",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
	})
}

// PostAdminLogout handles logout
func PostAdminLogout(c *fiber.Ctx) error {
	sess := middleware.GetSession(c)
	if err := sess.Destroy(); err != nil {
		return c.Redirect("/admin")
	}
	return c.Redirect("/admin/login")
}

// GetAdminDashboard renders the admin dashboard
func GetAdminDashboard(c *fiber.Ctx) error {
	sess := middleware.GetSession(c)

	user := map[string]interface{}{
		"email":       sess.Get("email"),
		"displayName": sess.Get("displayName"),
		"uid":         sess.Get("uid"),
	}

	return c.Render("admin/dashboard", fiber.Map{
		"title": "Admin Dashboard",
		"user":  user,
	})
}
