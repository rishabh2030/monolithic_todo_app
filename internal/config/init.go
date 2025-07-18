package config

import (
	"os"

	"gorm.io/gorm"
)

type Config struct {
	Dsn       string
	DbEngine  string
	Db        *gorm.DB
	Migration bool
}

func LoadConfig() *Config {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		url = "user=postgres password=postgres dbname=todo-app port=5432 sslmode=disable TimeZone=Asia/Kolkata host=localhost"
	}
	engine := os.Getenv("DATABASE_ENGINE")
	if engine == "" {
		engine = "postgres"
	}
	migration := os.Getenv("MIGRATION")
	if migration == "" {
		migration = "true"
	}
	return &Config{
		Dsn:       url,
		DbEngine:  engine,
		Migration: migration == "true",
	}
}
