package models

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
)

type Task struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
}

func (t Task) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Description, validation.Required, validation.Length(1, 140)),
	)
}
