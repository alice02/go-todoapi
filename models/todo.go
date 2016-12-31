package models

import (
	"time"
)

type Todo struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
	Description string     `json:"description" valid:"required,alphanum"`
	Completed   bool       `json:"completed"`
}
