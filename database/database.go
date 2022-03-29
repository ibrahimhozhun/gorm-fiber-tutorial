package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func Open() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("book.db"))

	// Migration
	db.AutoMigrate(&Book{})

	return db, err
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()

	// Close database connection
	sqlDB.Close()

	return err
}
