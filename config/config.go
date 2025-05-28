package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found (OK in production)")
    }
}

var (
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
	PORT 	   string
)

func LoadConfig() {
    DBHost = os.Getenv("DB_HOST")
    DBUser = os.Getenv("DB_USER")
    DBPassword = os.Getenv("DB_PASSWORD")
    DBName = os.Getenv("DB_NAME")
    DBPort = os.Getenv("DB_PORT")
	PORT = os.Getenv("PORT")
}
