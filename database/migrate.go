package database

import (
	"fmt"
)

var models []interface{}

func RegisterModel(m ...interface{}) {
	models = append(models, m...)
}

func Migrate() error {
	db := GetDB()

	// Example migration
	err := db.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	return nil
}
