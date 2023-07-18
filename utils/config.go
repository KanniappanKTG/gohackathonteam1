package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel   string
	DBHost     string
	DBPort     string
	DBName     string
	DBSchema   string
	DBUserName string
	DBPassword string
}

var Cfg = &Config{}

func init() {
	GetConfig()
}

func GetConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
		return err
	}
	Cfg.LogLevel = os.Getenv("LOGLEVEL")
	Cfg.DBHost = os.Getenv("DB_HOST")
	Cfg.DBPort = os.Getenv("DB_PORT")
	Cfg.DBName = os.Getenv("DB_DATABASE")
	Cfg.DBSchema = os.Getenv("DB_SCHEMA")
	Cfg.DBUserName = os.Getenv("DB_USER")
	Cfg.DBPassword = os.Getenv("DB_PASSWORD")
	return nil
}
