package repository

import "reservation/interanal/models"

type BokingRepo interface {
	List() []models.Reservation
	GetReservationDetails(id uint) (models.ReservationDetails, error)
	Create(reservation models.Reservation) error
	Update(reservation models.ReservationDetails) error
	Delete(id uint) error
}