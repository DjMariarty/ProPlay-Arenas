package config

import (
	"fmt"
	"os"
	"venue-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
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
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %w", err)
	}

	// Проверяем существование старой колонки type и переименовываем при необходимости
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения sql.DB: %w", err)
	}

	var exists bool
	err = sqlDB.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM information_schema.columns 
			WHERE table_name = 'venues' AND column_name = 'type'
		)
	`).Scan(&exists)

	if err == nil && exists {
		// Переименовываем старую колонку type в venue_type
		_, err = sqlDB.Exec("ALTER TABLE venues RENAME COLUMN type TO venue_type")
		if err != nil {
			return nil, fmt.Errorf("ошибка переименования колонки: %w", err)
		}
	}

	if err := db.AutoMigrate(&models.Venue{}); err != nil {
		return nil, fmt.Errorf("ошибка при миграции базы данных: %w", err)
	}

	return db, nil
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
