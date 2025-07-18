package migrations

import (
	"todo/internal/config"
	"todo/internal/models"
)

func CreateMigrations(config *config.Config) {
	if config.Migration {
		db := config.Db

		// Enable uuid-ossp extension
		if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
			panic("failed to create uuid-ossp extension: " + err.Error())
		}

		if err := db.AutoMigrate(&models.User{}); err != nil {
			panic("failed to migrate User model: " + err.Error())
		}
		if err := db.AutoMigrate(&models.Todos{}); err != nil {
			panic("failed to migrate Todos model: " + err.Error())
		}
	}
}
