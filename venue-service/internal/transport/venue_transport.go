package transport

import (
	"log/slog"
	"net/http"
	"strconv"
	"venue-service/internal/models"
	"venue-service/internal/services"

	"github.com/gin-gonic/gin"
)

type VenueHandler struct {
	service services.VenueService
	logger  *slog.Logger
}

func NewVenueHandler(service services.VenueService, logger *slog.Logger) *VenueHandler {
	return &VenueHandler{
		service: service,
		logger:  logger.With("layer", "transport"),
	}
}

func (h *VenueHandler) RegisterRoutes(r *gin.Engine) {
	venue := r.Group("/venue")
	{
		venue.POST("/", h.Create)
		venue.GET("/:id", h.GetByID)
	}
}

func (h *VenueHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Неверный формат ID", "id", idStr, "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный формат ID",
		})
		return
	}

	venue, err := h.service.GetByID(uint(id))
	if err != nil {
		h.logger.Error("Ошибка получения площадки по ID", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.logger.Info("Площадка успешно получена", "id", id)
	c.JSON(http.StatusOK, venue)
}

func (h *VenueHandler) Create(c *gin.Context) {
	var venue models.Venue
	if err := c.ShouldBindJSON(&venue); err != nil {
		h.logger.Error("Ошибка парсинга JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Create(&venue); err != nil {
		h.logger.Error("Ошибка создания площадки", "venue_type", venue.VenueType, "owner_id", venue.OwnerID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.logger.Info("Площадка успешно создана", "id", venue.ID, "venue_type", venue.VenueType, "owner_id", venue.OwnerID)
	c.JSON(http.StatusCreated, venue)
}
