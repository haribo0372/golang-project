package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() EnvModel {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	dbuser := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	appId := os.Getenv("APP_ID")
	return EnvModel{dbuser, dbpassword, dbname, appId}
}

type EnvModel struct {
	DbUser     string
	DbPassword string
	DbName     string
	AppId      string
}
