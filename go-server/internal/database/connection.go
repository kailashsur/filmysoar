package database

import (
	"fmt"
	"log"
	"time"

	"filmyfly-go-fiber/internal/config"
	"filmyfly-go-fiber/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establishes a connection to the PostgreSQL database
func Connect(cfg *config.Config) error {
	var err error

	// Configure GORM logger
	gormLogger := logger.Default
	if cfg.Environment == "production" {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else if cfg.PrismaLogQueries {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	// Open database connection
	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying SQL database
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Migrate before serving requests so fresh deployments do not require
	// undocumented manual schema setup.
	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("failed to migrate database schema: %w", err)
	}

	log.Println("✅ Database connection established")

	return nil
}

// AutoMigrate runs database migrations for all models
func AutoMigrate() error {
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Movie{},
		&models.TrendingMovie{},
		&models.Setting{},
		&models.StaticPage{},
		&models.AstroSetting{},
	); err != nil {
		return err
	}

	// Enable PostgreSQL's trigram operator class for fast substring searches.
	// This is idempotent and is required by the title GIN index below.
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS pg_trgm`).Error; err != nil {
		return fmt.Errorf("failed to enable pg_trgm extension: %w", err)
	}

	// These indexes cover the API's common filters, sorts, joins, and lookups.
	// IF NOT EXISTS keeps this migration safe to run on every startup.
	indexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_movies_slug ON movies (slug)`,
		`CREATE INDEX IF NOT EXISTS idx_movies_created_at ON movies ("createdAt" DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_movies_category_id ON movies ("categoryId")`,
		`CREATE INDEX IF NOT EXISTS idx_movies_release_year ON movies ("releaseYear")`,
		`CREATE INDEX IF NOT EXISTS idx_movies_title_lower ON movies (LOWER(title))`,
		`CREATE INDEX IF NOT EXISTS idx_movies_title_trgm ON movies USING gin (LOWER(title) gin_trgm_ops)`,
		`CREATE INDEX IF NOT EXISTS idx_static_pages_slug_published ON static_pages (slug, "is_published")`,
		`CREATE INDEX IF NOT EXISTS idx_trending_movies_order ON trending_movies ("order")`,
	}
	for _, statement := range indexes {
		if err := DB.Exec(statement).Error; err != nil {
			return err
		}
	}

	return nil
}
