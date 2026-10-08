package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	AppEnv     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	// Try loading .env from current directory or root
	_ = godotenv.Load()
	_ = godotenv.Load("../.env")

	return &Config{
		Port:       getEnv("APP_PORT", "8080"),
		AppEnv:     getEnv("APP_ENV", "development"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "appuser"),
		DBPassword: getEnv("DB_PASSWORD", "appsecret"),
		DBName:     getEnv("DB_NAME", "appdb"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
