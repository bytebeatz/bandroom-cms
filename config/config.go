package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	Env         string
	DBUrl       string
	GCSBucket   string
	GCSEnabled  bool
	GoogleCreds string
	JWTSecret   string
}

var AppConfig *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found or failed to load (continuing with OS env)")
	}

	viper.SetConfigFile("config/env.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No env.yaml found, relying on ENV vars")
	}

	AppConfig = &Config{
		Port:        getString("PORT", "8080"),
		Env:         getString("ENV", "development"),
		DBUrl:       getString("DATABASE_URL", ""),
		GCSBucket:   getString("GCS_BUCKET", ""),
		GCSEnabled:  viper.GetBool("GCS_ENABLED"),
		GoogleCreds: getString("GOOGLE_APPLICATION_CREDENTIALS", ""),
		JWTSecret:   getString("JWT_SECRET", "your-secret-here"),
	}

	log.Printf("Loaded DATABASE_URL: %s", AppConfig.DBUrl)
	log.Printf("Loaded JWT_SECRET: %s", AppConfig.JWTSecret)
}

func getString(key string, fallback string) string {
	if val := viper.GetString(key); val != "" {
		return val
	}
	return fallback
}
