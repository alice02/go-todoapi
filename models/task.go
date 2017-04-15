package models

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type TaskModelInterface interface {
	FindAll() ([]Task, error)
	FindByID(id int) (*Task, error)
	Save(model *Task) error
	Update(model *Task) error
	Delete(model *Task) error
}

type taskModel struct {
	session *gorm.DB
}

type Task struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
}

func NewTaskModel(db *gorm.DB) *taskModel {
	return &taskModel{db}
}

func (t Task) Validate() error {
	return validation.ValidateStruct(
		&t,
		validation.Field(
			&t.Description,
			validation.Required,
			validation.Length(1, 140)),
	)
}

func (t taskModel) FindAll() ([]Task, error) {
	tasks := []Task{}
	err := t.session.Find(&tasks)
	if err.Error != nil {
		return nil, err.Error
	}
	return tasks, nil
}

func (t taskModel) FindByID(id int) (*Task, error) {
	task := &Task{}
	err := t.session.First(task, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return task, nil
}

func (t taskModel) Save(model *Task) error {
	err := t.session.Create(model)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (t taskModel) Update(model *Task) error {
	err := t.session.Save(model)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (t taskModel) Delete(model *Task) error {
	err := t.session.Delete(model)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
