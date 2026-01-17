package services

import (
	"errors"
	"log/slog"
	"venue-service/internal/models"
	"venue-service/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrVenueNotFound = errors.New("venue not found")
)

type VenueService interface {
	GetByID(id uint) (*models.Venue, error)
	// GetList(filter VenueFilter) ([]models.Venue, error)
	Create(venue *models.Venue) error
	// Update(venue *models.Venue) error
	// Delete(id uint) error
}

type venueService struct {
	repository repository.VenueRepository
	logger     *slog.Logger
}

func NewVenueService(repository repository.VenueRepository, logger *slog.Logger) VenueService {
	return &venueService{
		repository: repository,
		logger:     logger.With("layer", "service"),
	}
}

func (s *venueService) GetByID(id uint) (*models.Venue, error) {
	venue, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrVenueNotFound
		}
		s.logger.Error("Ошибка получения площадки по ID", "id", id, "error", err)
		return nil, err
	}

	return venue, nil
}

func (s *venueService) Create(v *models.Venue) error {

	if err := s.repository.Create(v); err != nil {
		s.logger.Error("Ошибка создания площадки", "venue_type", v.VenueType, "owner_id", v.OwnerID, "error", err)
		return err
	}

	return nil
}
