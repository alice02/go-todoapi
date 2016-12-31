package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alice02/go_react_todoapp/database"
	"github.com/alice02/go_react_todoapp/models"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"message,omitempty"`
}

type Data struct {
	Todos []models.Todo `json:"tasks,omitempty"`
	Todo  *models.Todo  `json:"task,omitempty"`
	Info  string        `json:"info,omitempty"`
}

// GET /api/todo
func GetTodos(c echo.Context) error {
	var todos []models.Todo

	db := database.GetDb()
	db.Find(&todos)

	fmt.Println(todos)

	data := Data{
		Todos: todos,
	}
	fmt.Println(data)
	response := ResponseMessage{
		Status: "success",
		Data:   data,
	}

	return c.JSON(http.StatusOK, response)
}

// GET /api/todo/:id
func GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := fetchTodoById(id)
	if err != nil {
		data := Data{
			Info: err.Error(),
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	data := Data{
		Todo: &todo,
	}
	response := ResponseMessage{
		Status: "success",
		Data:   data,
	}

	return c.JSON(http.StatusOK, response)
}

// POST /api/todo
func PostTodo(c echo.Context) error {
	db := database.GetDb()

	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		response := ResponseMessage{}
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err := govalidator.ValidateStruct(todo)
	if err != nil {
		response := ResponseMessage{}
		return c.JSON(http.StatusBadRequest, response)
	}

	db.Create(&todo)

	return c.JSON(http.StatusOK, todo)
}

// PUT /api/todo/:id
func PutTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, _ := fetchTodoById(id)

	if err := c.Bind(&todo); err != nil {
		response := ResponseMessage{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	db := database.GetDb()
	db.Save(&todo)

	return c.JSON(http.StatusOK, todo)
}

// DELETE /api/todo/:id
func DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, _ := fetchTodoById(id)

	db := database.GetDb()
	db.Delete(&todo)

	return c.JSON(http.StatusOK, nil)
}

func fetchTodoById(id int) (models.Todo, error) {
	var todo models.Todo

	db := database.GetDb()
	if err := db.First(&todo, id).Error; err != nil {
		return todo, err
	}

	return todo, nil
}
