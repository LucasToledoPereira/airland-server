package itinerary_handler

import (
	"airland-server/src/domain/entities"
	requests "airland-server/src/domain/itinerary/commands/requests"
	"airland-server/src/domain/repository"
)

type ItineraryHandler struct {
	repository repository.IItineraryRepository
}

func NewItineraryHandler(r repository.IItineraryRepository) (handler *ItineraryHandler) {
	return &ItineraryHandler{
		repository: r,
	}
}

func (h *ItineraryHandler) Create(command requests.CreateItineraryCommandRequest) (it entities.Itinerary, err error) {
	it = command.ToEntity()
	err = h.repository.Create(&it)
	return it, err
}
