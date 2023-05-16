package itinerary_commands

import (
	"airland-server/src/domain/entities"
	"time"

	"github.com/google/uuid"
)

type CreateItineraryCommandRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Budget      float32   `json:"budget"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	UserID      string
}

func (c *CreateItineraryCommandRequest) ToEntity() (itinerary entities.Itinerary) {
	return entities.Itinerary{
		ID:          uuid.New(),
		Title:       c.Title,
		Description: c.Description,
		Budget:      c.Budget,
		StartAt:     c.StartAt,
		EndAt:       c.EndAt,
	}
}
