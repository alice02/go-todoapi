package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Description string `json:"description"`
	Status bool `json:"status"`
}
