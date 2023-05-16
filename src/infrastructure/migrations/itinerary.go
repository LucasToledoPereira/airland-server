package migrations

import (
	"airland-server/src/domain/entities"

	"gorm.io/gorm"
)

func migrateItinerary(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&entities.Itinerary{}) {
		//db.Migrator().DropTable(&entities.Itinerary{})
		return nil
	}

	return db.Migrator().CreateTable(&entities.Itinerary{})
}
