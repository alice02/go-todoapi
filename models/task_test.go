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

func TestSaveAndFind(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Task{})
	u := NewTaskModel(db)
	task := Task{
		Description: "test",
		Completed:   false,
	}
	err = u.Save(&task)
	if err != nil {
		t.Errorf("database save failed")
	}
	db.DropTableIfExists(&Task{})
}

func TestUpdate(t *testing.T) {
}
