package datastore

import (
	"fmt"

	"airland-server/src/cross_cutting/config"
	"airland-server/src/infrastructure/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Data struct {
	DB *gorm.DB
}

func NewData() (data *Data, err error) {
	db, err := connectDB()

	if err != nil {
		return nil, err
	}

	/*if err := prepareDB(db); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}*/

	if err := migrateDB(db); err != nil {
		return nil, err
	}

	return &Data{
		DB: db,
	}, nil
}

func connectDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.User,
		config.C.Database.Password,
		config.C.Database.DBName,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(db *gorm.DB) (err error) {
	if err := migrations.Migrate(db); err != nil {
		return err
	}

	/*if err := db.AutoMigrate(&entities.User{}); err != nil {
		return err
	}*/

	return nil
}

/*
func prepareDB(db *gorm.DB) (err error) {
	if err := db.AutoMigrate(&entities.CustomMigrations{}); err != nil {
		return err
	}

	files, _ := filepath.Glob(utils.Path() + "/migrations/*.sql")

	for stix := range files {
		slice := strings.SplitN(filepath.Base(files[stix]), ".", 2)
		id, err := strconv.Atoi(slice[0])

		if err != nil {
			return err
		}

		var migration entities.CustomMigrations
		result := db.Where(entities.CustomMigrations{ID: id}).First(&migration)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			content, _ := os.ReadFile(files[stix])
			result = db.Exec(string(content))

			fmt.Println(string(content))

			if result.Error != nil {
				return result.Error
			}

			migration.ID = id
			migration.Name = slice[1]
			result = db.Create(&migration)

			if result.Error != nil {
				return result.Error
			}
		}

	}

	return nil
}*/
