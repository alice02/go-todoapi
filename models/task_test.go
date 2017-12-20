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

func TestValidateWithInvalidDescription(t *testing.T) {
	emptyDescription := Task{
		Description: "",
		Completed:   false,
	}
	actual := emptyDescription.Validate()
	if actual == nil {
		t.Errorf("got nil want validate error")
	}
	expected := "description: cannot be blank."
	if actual.Error() != expected {
		t.Errorf("got %v want %v", actual, expected)
	}

	overLengthDescription := Task{
		Description: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Completed:   false,
	}
	actual = overLengthDescription.Validate()
	if actual == nil {
		t.Errorf("got nil want validate error")
	}
	expected = "description: the length must be between 1 and 140."
	if actual.Error() != expected {
		t.Errorf("got %v want %v", actual, expected)
	}

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
	db.DropTableIfExists(&Task{})
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
		t.Errorf("got length %v want %v", len(actual), len(expected))
	}
	for i := range actual {
		if actual[i].Description != expected[i].Description {
			t.Errorf("got %v want %v", actual[i].Description, expected[i].Description)
		}
		if actual[i].Completed != expected[i].Completed {
			t.Errorf("got %v want %v", actual[i].Completed, expected[i].Completed)
		}
	}

	db.DropTableIfExists(&Task{})
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
	db.DropTableIfExists(&Task{})
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

func TestFindById(t *testing.T) {
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
	db.DropTableIfExists(&Task{})
	db.AutoMigrate(&Task{})
	u := NewTaskModel(db)

	for _, task := range testData {
		err := u.Save(&task)
		if err != nil {
			t.Errorf("save failed")
		}
		actual, _ := u.FindByID(int(task.ID))
		if actual.Description != task.Description {
			t.Errorf("got %v want %v", actual.Description, task.Description)
		}
		if actual.Completed != task.Completed {
			t.Errorf("got %v want %v", actual.Completed, task.Completed)
		}
	}
	db.DropTableIfExists(&Task{})
}
