package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL      string
	RedisAddr        string
	APIPort          string
	JWTSecret        string
	OTPExpiryMinutes string
}

func Load() *Config {
	cfg := &Config{
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		RedisAddr:        os.Getenv("REDIS_ADDR"),
		APIPort:          os.Getenv("API_PORT"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		OTPExpiryMinutes: os.Getenv("OTP_EXPIRY_MINUTES"),
	}

	if cfg.DatabaseURL == "" || cfg.JWTSecret == "" {
		log.Fatal("Missing required environment variables")
	}

	if cfg.APIPort == "" {
		cfg.APIPort = "8080"
	}

	return cfg
}
