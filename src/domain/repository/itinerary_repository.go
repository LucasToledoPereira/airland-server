package repository

import (
	"airland-server/src/domain/entities"

	"github.com/google/uuid"
)

type IItineraryRepository interface {
	GetById(id uuid.UUID, itinerary *entities.Itinerary) (err error)
	GetAll(userId string, itineraries *[]entities.Itinerary) (err error)
	Create(itinerary *entities.Itinerary) (err error)
}
