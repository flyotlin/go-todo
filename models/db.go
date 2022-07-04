package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	dbURL = "test.db"
)

func openDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}

	db.AutoMigrate(&Todo{})

	return db
}
