package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	BotToken string
	BaseURL  string
}

var Env *EnvConfig

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Relying on system environment variables.")
	}

	Env = &EnvConfig{
		BotToken: getEnv("BOT_TOKEN", ""),
		BaseURL:  getEnv("BASE_URL", "https://api.dictionaryapi.dev/api/v2/entries/en/"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
