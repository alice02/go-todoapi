package models

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestNewTaskModel(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	actual := NewTaskModel(db)
	fmt.Println(reflect.TypeOf(actual))
}

func TestFindAll(t *testing.T) {

}
