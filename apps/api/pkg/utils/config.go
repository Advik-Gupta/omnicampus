package utils

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type cfg struct {
	Port             string `env:"PORT" envDefault:"8080"`
	PostgresHost     string `env:"POSTGRES_HOST,notEmpty"`
	PostgresPort     string `env:"POSTGRES_PORT,notEmpty"`
	PostgresUser     string `env:"POSTGRES_USER,notEmpty"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,notEmpty"`
	PostgresDB       string `env:"POSTGRES_DB,notEmpty"`
}

var Config cfg

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf("No .env file found")
	}

	if err := env.Parse(&Config); err != nil {
		fmt.Errorf("%+v", err)
	}

	fmt.Printf("Configuration successfully loaded\n")
}