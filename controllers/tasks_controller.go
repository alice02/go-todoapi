package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alice02/go-todoapi/database"
	"github.com/alice02/go-todoapi/models"
	"github.com/labstack/echo"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"message,omitempty"`
}

type Data struct {
	Tasks []models.Task `json:"tasks,omitempty"`
	Task  *models.Task  `json:"task,omitempty"`
	Info  string        `json:"info,omitempty"`
}

// GET /api/task
func GetTasks(c echo.Context) error {
	var tasks []models.Task

	db := database.GetDb()
	db.Find(&tasks)

	data := Data{
		Tasks: tasks,
	}
	response := ResponseMessage{
		Status: "success",
		Data:   data,
	}

	return c.JSON(http.StatusOK, response)
}

// GET /api/task/:id
func GetTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := fetchTaskById(id)
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
		Task: task,
	}
	response := ResponseMessage{
		Status: "success",
		Data:   data,
	}

	return c.JSON(http.StatusOK, response)
}

// POST /api/task
func PostTask(c echo.Context) error {
	db := database.GetDb()

	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		data := Data{
			Info: "invalid request body",
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := task.Validate(); err != nil {
		data := Data{
			Info: err.Error(),
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	db.Create(&task)

	return c.JSON(http.StatusOK, task)
}

// PUT /api/task/:id
func PutTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := fetchTaskById(id)
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

	if err := c.Bind(task); err != nil {
		fmt.Println(err)
		data := Data{
			Info: "invalid request body",
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := task.Validate(); err != nil {
		data := Data{
			Info: err.Error(),
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	db := database.GetDb()
	db.Save(task)

	return c.JSON(http.StatusOK, task)
}

// DELETE /api/task/:id
func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, _ := fetchTaskById(id)

	db := database.GetDb()
	db.Delete(&task)

	return c.JSON(http.StatusOK, nil)
}

func fetchTaskById(id int) (*models.Task, error) {
	var task models.Task

	db := database.GetDb()
	if err := db.First(&task, id).Error; err != nil {
		return &task, err
	}

	return &task, nil
}
