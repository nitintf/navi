package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// oncer for the config loading
var loadOnce sync.Once

func get(name, def string) string {
	// Load .env file only once
	loadOnce.Do(func() {
		godotenv.Load(".env")
	})

	if value := os.Getenv(name); value != "" {
		return value
	}

	return def
}

var (
	GeminiKey   = get("GEMINI_KEY", "")
	Environment = get("GO_ENV", "development")
)
