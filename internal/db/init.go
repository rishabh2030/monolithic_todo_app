package db

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDB(database_engine string, dsn string) (*gorm.DB, error) {
	switch database_engine {
	case "postgres":
		return InitPostgresDB(dsn), nil
	default:
		return nil, fmt.Errorf("unsupported database engine: %s", database_engine)
	}
}
