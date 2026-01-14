package config

import (
	"fmt"
	"os"
	"venue-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", "postgres"),
		GetEnv("DB_NAME", "venue_db"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_SSLMODE", "disable"),
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return fmt.Errorf("Ошибка при подключении к базе данных: %w", err)
	}

	DB = db

	// Автомиграция моделей
	if err := DB.AutoMigrate(&models.Venue{}); err != nil {
		return fmt.Errorf("Ошибка при миграции базы данных: %w", err)
	}

	return nil
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
