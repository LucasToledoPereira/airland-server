package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Itinerary struct {
	ID          uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Name        string    `gorm:"not null; default:null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Itinerary) TableName() string { return "itinerary" }
