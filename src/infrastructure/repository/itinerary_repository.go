package repository

import (
	"airland-server/src/domain/entities"
	"airland-server/src/infrastructure/datastore"

	"github.com/google/uuid"
)

type ItineraryRepo struct {
	data *datastore.Data
}

func NewItineraryRepo(d *datastore.Data) (itr *ItineraryRepo) {
	return &ItineraryRepo{
		data: d,
	}
}

func (ir ItineraryRepo) GetById(id uuid.UUID, itinerary *entities.Itinerary) (err error) {
	return nil
}

func (ir ItineraryRepo) GetAll(userId string, itineraries *[]entities.Itinerary) (err error) {
	return nil
}

func (ir ItineraryRepo) Create(itinerary *entities.Itinerary) (err error) {
	err = ir.data.DB.Create(&itinerary).Error
	return err
}
