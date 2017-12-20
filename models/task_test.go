package models

import (
	"reflect"
	"testing"

	"github.com/alice02/go-todoapi/database"
)

func TestNewTaskModel(t *testing.T) {
	u := NewTaskModel(nil)
	expected := "*models.taskModel"
	actual := reflect.TypeOf(u).String()
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestValidate(t *testing.T) {

}

func TestValidateWithInvalidDescription(t *testing.T) {

}

func TestSaveAndFind(t *testing.T) {
	expected := []Task{
		{
			Description: "test1",
			Completed:   false,
		},
		{
			Description: "test2",
			Completed:   true,
		},
		{
			Description: "",
			Completed:   false,
		},
	}

	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Task{})
	u := NewTaskModel(db)

	for _, task := range expected {
		err = u.Save(&task)
		if err != nil {
			t.Errorf("database save failed")
		}
	}
	actual, err := u.FindAll()

	if len(actual) != len(expected) {
		t.Errorf("got length %v want %v", len(actual), (expected))
	}
	for i := range actual {
		if actual[i].Description != expected[i].Description {
			t.Errorf("got %v want %v", actual[i].Description, expected[i].Description)
		}
		if actual[i].Completed != expected[i].Completed {
			t.Errorf("got %v want %v", actual[i].Description, expected[i].Description)
		}
	}

	db.DropTableIfExists(&Task{})
}

func TestFindByID(t *testing.T) {

}

func TestUpdate(t *testing.T) {
	testData := []Task{
		{
			Description: "test1",
			Completed:   false,
		},
		{
			Description: "test2",
			Completed:   true,
		},
		{
			Description: "",
			Completed:   false,
		},
	}

	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Task{})
	u := NewTaskModel(db)

	for _, task := range testData {
		if err = u.Save(&task); err != nil {
			t.Errorf("database save failed")
		}
		task.Description = "updated"
		if err = u.Update(&task); err != nil {
			t.Errorf("database save failed")
		}
		updatedTask, err := u.FindByID(int(task.ID))
		if err != nil {
			t.Errorf("database select failed")
		}
		if task.Description != updatedTask.Description {
			t.Errorf("got %v want %v", "updated", task.Description)
		}

	}
	db.DropTableIfExists(&Task{})
}

func TestDelete(t *testing.T) {

}
