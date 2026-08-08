package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	logDir       string
	logFile      string
	errorLogFile string
}

var AppLogger *Logger

// InitLogger initializes the logger
func InitLogger() *Logger {
	logDir := "logs"
	logFile := filepath.Join(logDir, "app.log")
	errorLogFile := filepath.Join(logDir, "error.log")

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
	}

	logger := &Logger{
		logDir:       logDir,
		logFile:      logFile,
		errorLogFile: errorLogFile,
	}

	AppLogger = logger
	return logger
}

func (l *Logger) formatLog(level, message string, args ...interface{}) string {
	timestamp := time.Now().Format(time.RFC3339)
	formattedMessage := fmt.Sprintf(message, args...)
	return fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, formattedMessage)
}

func (l *Logger) writeLog(file, message string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return // Silently fail
	}
	defer f.Close()

	f.WriteString(message)
}

func (l *Logger) Info(message string, args ...interface{}) {
	logMessage := l.formatLog("INFO", message, args...)
	l.writeLog(l.logFile, logMessage)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	logMessage := l.formatLog("WARN", message, args...)
	l.writeLog(l.logFile, logMessage)
	l.writeLog(l.errorLogFile, logMessage)
}

func (l *Logger) Error(message string, args ...interface{}) {
	logMessage := l.formatLog("ERROR", message, args...)
	l.writeLog(l.logFile, logMessage)
	l.writeLog(l.errorLogFile, logMessage)
}

// Global logging functions
func Info(message string, args ...interface{}) {
	if AppLogger != nil {
		AppLogger.Info(message, args...)
	}
}

func Warn(message string, args ...interface{}) {
	if AppLogger != nil {
		AppLogger.Warn(message, args...)
	}
}

func Error(message string, args ...interface{}) {
	if AppLogger != nil {
		AppLogger.Error(message, args...)
	}
}
