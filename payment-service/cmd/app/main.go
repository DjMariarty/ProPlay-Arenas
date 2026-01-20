package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"payment-service/internal/config"
	"payment-service/internal/models"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("файл .env не загружен", "error", err)
	}

	db := config.ConnectDB()

	if err := db.AutoMigrate(
		&models.Payment{},
		&models.Refund{},
	); err != nil {
		slog.Error("ошибка миграции базы данных", "error", err)
		os.Exit(1)
	}

	slog.Info("миграция базы данных завершена")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","service":"payment-service"}`))
	})

	port := config.GetEnv("PORT", "8080")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	slog.Info("HTTP сервер запущен", "port", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("ошибка запуска HTTP сервера", "error", err)
		os.Exit(1)
	}
}
