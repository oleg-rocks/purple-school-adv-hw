package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Username string
	Password string
	Host string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	return &Config{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Host: os.Getenv("HOST"),
	}
}