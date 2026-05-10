package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBURL      string

	JWTAccessSecret  string
	JWTRefreshSecret string
	JWTAccessExpMin  int
	JWTRefreshExpDay int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	return &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "learnsphere"),
		DBURL:      getEnv("DATABASE_URL", ""),

		JWTAccessSecret:  getEnv("JWT_ACCESS_SECRET", "change-me-access"),
		JWTRefreshSecret: getEnv("JWT_REFRESH_SECRET", "change-me-refresh"),
		JWTAccessExpMin:  getEnvInt("JWT_ACCESS_EXP_MIN", 15),
		JWTRefreshExpDay: getEnvInt("JWT_REFRESH_EXP_DAY", 7),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}