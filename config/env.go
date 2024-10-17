package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string
}

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "3000"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PWD", ""),
		DBName:     getEnv("DB_NAME", "ecom"),
		DBAddress:  getEnv("DB_ADDR", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
