package migrations

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) (err error) {
	err = migrateItinerary(db)
	return err
}
