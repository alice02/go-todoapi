package database

import (
	"github.com/alice02/go-todoapi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "todo.db")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{})

	return db, nil
}
