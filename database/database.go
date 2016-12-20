package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/alice02/go_react_todoapp/models"
)

var (
	db *gorm.DB
	err error
)

func InitDb() {
	// connect to database
	db, err = gorm.Open("sqlite3", "todo.db")
	if err != nil {
		panic("failed to connect database")
	}

	// migration
	db.AutoMigrate(&models.Todo{})
}

func GetDb() *gorm.DB{
	return db
}
