package main

import (
	"fmt"
	"log"
	"net/http"

	"venue-service/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("ConnectDB: %v", err)
	}

	r := gin.Default()

	// Отключаем доверие прокси для локальной разработки
	// Для production используйте: r.SetTrustedProxies([]string{"127.0.0.1"})
	// TODO: не забыть включить доверие прокси для production
	r.SetTrustedProxies(nil)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.Run(fmt.Sprintf(":%s", config.GetEnv("PORT", "8080")))
}
