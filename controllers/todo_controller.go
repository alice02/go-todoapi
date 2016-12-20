package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/alice02/go_react_todoapp/models"
)

// GET /api/todo
func GetTodos(c echo.Context) error {
	todos := &models.Todo {
		Description: "hogehoge",
		Status: false,
	}
	return c.JSON(http.StatusOK, todos)
}

// GET /api/todo/:id
func GetTodo(c echo.Context) error {
	todos := &models.Todo {
		Description: "hogehoge",
		Status: false,
	}
	return c.JSON(http.StatusOK, todos)
}

// POST /api/todo
func PostTodo(c echo.Context) error {
	todos := &models.Todo {
		Description: "hogehoge",
		Status: false,
	}
	return c.JSON(http.StatusOK, todos)
}

// PUT /api/todo/:id
func PutTodo(c echo.Context) error {
	todos := &models.Todo {
		Description: "hogehoge",
		Status: false,
	}
	return c.JSON(http.StatusOK, todos)
}

// DELETE /api/todo/:id
func DeleteTodo(c echo.Context) error {
	todos := &models.Todo {
		Description: "hogehoge",
		Status: false,
	}
	return c.JSON(http.StatusOK, todos)
}
