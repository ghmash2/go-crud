package config

import (
	"os"
	"log"
)

type Config struct {
	AppPort string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func LoadConfig() Config {
    config := Config{
        AppPort: getEnv("APP_PORT"),
        DBHost: getEnv("DB_HOST"),
        DBPort: getEnv("DB_PORT"),
        DBName: getEnv("DB_NAME"),
        DBUser: getEnv("DB_USER"),
        DBPass: getEnv("DB_PASS"),
    }
	if config.DBName == "" || config.DBUser == "" || config.DBPass == "" {
		log.Fatalf("Environment variable DB_NAME / DB_USER / DB_PASS is not set")
	}
    return config
}