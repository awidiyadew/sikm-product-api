package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetenvInt(envName string) int {
	strEnv := os.Getenv(envName)
	i, err := strconv.Atoi(strEnv)
	if err != nil {
		log.Fatalf("error parsing env %s to int", envName)
	}
	return i
}

var (
	AppPort    = os.Getenv("APP_PORT")
	DBHost     = os.Getenv("DB_HOST")
	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
	DBPort     = GetenvInt("DB_PORT")
	JWTSecret  = os.Getenv("JWT_SECRET")
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
