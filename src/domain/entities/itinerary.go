package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Itinerary struct {
	ID          uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	UserID      string    `gorm:"not null; index;"`
	Title       string    `gorm:"not null;"`
	Description string
	Budget      float32
	StartAt     time.Time
	EndAt       time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Itinerary) TableName() string { return "itinerary" }
