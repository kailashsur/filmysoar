package admin

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

// GetLogs displays logs viewer page
func GetLogs(c *fiber.Ctx) error {
	return c.Render("admin/logs", fiber.Map{
		"title": "View Logs",
		"user":  getUserFromSession(c),
	})
}

// GetLogsData returns logs data as JSON
func GetLogsData(c *fiber.Ctx) error {
	appLogs, appSize := readLogFile("logs/app.log")
	errorLogs, errorSize := readLogFile("logs/error.log")

	return c.JSON(fiber.Map{
		"success":               true,
		"logs":                  appLogs,
		"errorLogs":             errorLogs,
		"logSize":               appSize,
		"errorLogSize":          errorSize,
		"logSizeFormatted":      formatBytes(appSize),
		"errorLogSizeFormatted": formatBytes(errorSize),
	})
}

// ClearLogs clears log files
func ClearLogs(c *fiber.Ctx) error {
	var req struct {
		Type string `json:"type"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request",
		})
	}

	var message string

	if req.Type == "all" || req.Type == "app" {
		os.WriteFile("logs/app.log", []byte(""), 0644)
	}

	if req.Type == "all" || req.Type == "error" {
		os.WriteFile("logs/error.log", []byte(""), 0644)
	}

	switch req.Type {
	case "all":
		message = "All logs cleared successfully"
	case "app":
		message = "Application logs cleared successfully"
	case "error":
		message = "Error logs cleared successfully"
	default:
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid log type",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": message,
	})
}

// DownloadLogs downloads log file
func DownloadLogs(c *fiber.Ctx) error {
	logType := c.Query("type", "app")

	var filename, filepath string
	if logType == "error" {
		filename = "error.log"
		filepath = "logs/error.log"
	} else {
		filename = "app.log"
		filepath = "logs/app.log"
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "Log file not found",
		})
	}

	c.Set("Content-Type", "text/plain")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	return c.Send(content)
}

// readLogFile reads last 100KB of log file
func readLogFile(path string) (string, int64) {
	info, err := os.Stat(path)
	if err != nil {
		return "No logs available yet.", 0
	}

	size := info.Size()

	// Read last 100KB or entire file if smaller
	maxBytes := int64(100 * 1024)
	startPos := int64(0)
	if size > maxBytes {
		startPos = size - maxBytes
	}

	file, err := os.Open(path)
	if err != nil {
		return "Error reading log file.", size
	}
	defer file.Close()

	file.Seek(startPos, 0)
	content := make([]byte, size-startPos)
	file.Read(content)

	return string(content), size
}

// formatBytes formats bytes to human readable format
func formatBytes(bytes int64) string {
	if bytes == 0 {
		return "0 Bytes"
	}

	k := float64(1024)
	sizes := []string{"Bytes", "KB", "MB", "GB"}

	i := 0
	size := float64(bytes)
	for size >= k && i < len(sizes)-1 {
		size /= k
		i++
	}

	return fmt.Sprintf("%.2f %s", size, sizes[i])
}
