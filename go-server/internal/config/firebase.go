package config

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

// InitializeFirebase initializes Firebase Admin SDK
func InitializeFirebase() error {
	ctx := context.Background()
	cfg := Load()

	// Try to initialize with service account key
	if cfg.FirebaseServiceAccountKey != "" {
		// Check if it's a JSON string (starts with { or contains "type")
		trimmed := strings.TrimSpace(cfg.FirebaseServiceAccountKey)
		if strings.HasPrefix(trimmed, "{") || strings.Contains(trimmed, "\"type\"") {
			// It's JSON content - parse and use it
			log.Println("📝 Parsing Firebase service account key from JSON...")

			// Validate JSON
			var jsonCheck map[string]interface{}
			if err := json.Unmarshal([]byte(trimmed), &jsonCheck); err != nil {
				log.Printf("⚠️  Invalid JSON in FIREBASE_SERVICE_ACCOUNT_KEY: %v", err)
				return err
			}

			// Use the JSON credentials
			opt := option.WithCredentialsJSON([]byte(trimmed))
			app, err := firebase.NewApp(ctx, nil, opt)
			if err != nil {
				log.Printf("⚠️  Firebase initialization with JSON failed: %v", err)
				return err
			}

			client, err := app.Auth(ctx)
			if err != nil {
				log.Printf("⚠️  Firebase Auth client creation failed: %v", err)
				return err
			}

			FirebaseAuth = client
			log.Println("✅ Firebase initialized with JSON credentials")
			return nil
		}

		// It's a file path
		log.Printf("📁 Using Firebase service account key file: %s", cfg.FirebaseServiceAccountKey)
		opt := option.WithCredentialsFile(cfg.FirebaseServiceAccountKey)
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Printf("⚠️  Firebase initialization with file failed: %v", err)
			return err
		}

		client, err := app.Auth(ctx)
		if err != nil {
			log.Printf("⚠️  Firebase Auth client creation failed: %v", err)
			return err
		}

		FirebaseAuth = client
		log.Println("✅ Firebase initialized with service account key file")
		return nil
	}

	// Try with default credentials
	log.Println("🔍 Trying Firebase with default credentials...")
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Printf("⚠️  Firebase initialization failed: %v", err)
		log.Println("⚠️  Admin authentication will not work without Firebase")
		return err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("⚠️  Firebase Auth client creation failed: %v", err)
		return err
	}

	FirebaseAuth = client
	log.Println("✅ Firebase initialized with default credentials")
	return nil
}
