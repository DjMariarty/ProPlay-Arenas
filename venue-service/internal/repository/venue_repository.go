package repository

import (
	"log/slog"
	"venue-service/internal/models"

	"gorm.io/gorm"
)

type VenueFilter struct {
	District  string
	VenueType models.VenueType
	HourPrice int
}

type VenueRepository interface {
	GetByID(id uint) (*models.Venue, error)
	// GetList(filter VenueFilter) ([]models.Venue, error)
	Create(venue *models.Venue) error
	// Update(venue *models.Venue) error
	// Delete(id uint) error
}

type venueRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewVenueRepository(db *gorm.DB, logger *slog.Logger) VenueRepository {
	return &venueRepository{
		db:     db,
		logger: logger.With("layer", "repository"),
	}
}

func (r *venueRepository) Create(venue *models.Venue) error {
	if err := r.db.Create(venue).Error; err != nil {
		r.logger.Error("Ошибка создания площадки", "error", err)
		return err
	}
	return nil
}

func (r *venueRepository) GetByID(id uint) (*models.Venue, error) {
	var venue models.Venue
	if err := r.db.First(&venue, id).Error; err != nil {
		r.logger.Error("Ошибка получения площадки по ID", "id", id, "error", err)
		return nil, err
	}
	return &venue, nil
}

// func (r *venueRepository) GetList(filter VenueFilter) ([]models.Venue, error) {
// 	query := r.db.Model(&models.Venue{})
// 	if filter.District != "" {
// 		query = query.Where("district = ?", filter.District)
// 	}
// 	if filter.VenueType != "" {
// 		query = query.Where("venue_type = ?", filter.VenueType)
// 	}
// 	if filter.HourPrice > 0 {
// 		query = query.Where("hour_price = ?", filter.HourPrice)
// 	}

// 	var venues []models.Venue
// 	if err := query.Find(&venues).Error; err != nil {
// 		r.logger.Error("Ошибка получения всех площадок", "error", err)
// 		return nil, err
// 	}
// 	return venues, nil
// }

// func (r *venueRepository) Update(venue *models.Venue) error {
// 	if err := r.db.Save(venue).Error; err != nil {
// 		r.logger.Error("Ошибка обновления площадки", "error", err)
// 		return err
// 	}
// 	return nil
// }

// func (r *venueRepository) Delete(id uint) error {
// 	if err := r.db.Delete(&models.Venue{}, id).Error; err != nil {
// 		r.logger.Error("Ошибка удаления площадки", "id", id, "error", err)
// 		return err
// 	}
// 	return nil
// }
