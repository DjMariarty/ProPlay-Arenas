package main

import (
	"log/slog"
	"os"

	"gateway/internal/config"
	"gateway/internal/transport"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Загружаем .env только для локальной разработки
	// В Docker все переменные передаются через docker-compose.yaml
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			// Игнорируем ошибку, если .env файл не найден
		}
	}

	// Устанавливаем режим Gin
	if os.Getenv("ENV") == "production" || os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	jwtSecret := config.GetEnv("JWT_SECRET", "")
	if jwtSecret == "" {
		slog.Error("JWT_SECRET не задан")
		os.Exit(1)
	}

	cfg := transport.Config{
		JWTSecret:             jwtSecret,
		UserServiceURL:        config.GetEnv("USER_SERVICE_URL", "http://localhost:8080"),
		VenueServiceURL:       config.GetEnv("VENUE_SERVICE_URL", "http://localhost:8082"),
		ReservationServiceURL: config.GetEnv("RESERVATION_SERVICE_URL", "http://localhost:8081"),
		PaymentServiceURL:     config.GetEnv("PAYMENT_SERVICE_URL", "http://localhost:8084"),
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	if err := transport.RegisterRoutes(router, cfg); err != nil {
		slog.Error("ошибка настройки маршрутов", "error", err)
		os.Exit(1)
	}

	port := config.GetEnv("PORT", "8085")
	slog.Info("gateway запущен", "port", port)
	if err := router.Run(":" + port); err != nil {
		slog.Error("не удалось запустить сервер", "error", err)
		os.Exit(1)
	}
}
