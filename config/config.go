package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PushCallDetails string
}

var AppConfig *Config

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
	AppConfig = &Config{
		PushCallDetails: os.Getenv("pushCallDetails"),
	}
	log.Println("Loaded configuration:", AppConfig)
}
