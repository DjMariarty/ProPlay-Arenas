package transport

import (
	"log/slog"
	"venue-service/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	logger *slog.Logger,
	venueService services.VenueService,
) {
	venueHandler := NewVenueHandler(venueService, logger)
	venueHandler.RegisterRoutes(router)
}
