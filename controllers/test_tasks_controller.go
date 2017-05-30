package controllers

import (
	"testing"

	"github.com/alice02/go-todoapi/models"
)

type mockTaskModel struct {
}

func (t mockTaskModel) FindAll() ([]models.Task, error) {
	return nil, nil
}

func (t mockTaskModel) FindByID(id int) (*models.Task, error) {
	return nil, nil
}

func (t mockTaskModel) Save(model *models.Task) error {
	return nil
}

func (t mockTaskModel) Update(model *models.Task) error {
	return nil
}

func (t mockTaskModel) Delete(model *models.Task) error {
	return nil
}

func TestNewTaskController(t *testing.T) {
	taskModel := new(mockTaskModel)
	actual := NewTaskController(taskModel)
	print(actual)
}
