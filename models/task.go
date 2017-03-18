package models

import (
	"time"
)

type Task struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
}
