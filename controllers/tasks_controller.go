package controllers

import (
	"net/http"
	"strconv"

	"github.com/alice02/go-todoapi/models"
	"github.com/labstack/echo"
)

type (
	taskController struct {
		taskModel models.TaskModelInterface
	}
)

func NewTaskController(t models.TaskModelInterface) *taskController {
	return &taskController{t}
}

// GET /api/tasks
func (tc taskController) GetTasks(c echo.Context) error {
	tasks, err := tc.taskModel.FindAll()
	if err != nil {
		response := ResponseMessage{
			Status:  "failed",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := Data{
		Tasks: tasks,
	}
	response := ResponseMessage{
		Status: "success",
		Data:   data,
	}

	return c.JSON(http.StatusOK, response)
}

// GET /api/tasks/:id
func (tc taskController) GetTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := tc.taskModel.FindByID(id)
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

// POST /api/tasks
func (tc taskController) PostTask(c echo.Context) error {
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

	err := tc.taskModel.Save(task)
	if err != nil {
		data := Data{
			Info: err.Error(),
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusInternalServerError, response)
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

// PUT /api/tasks/:id
func (tc taskController) PutTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := tc.taskModel.FindByID(id)
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

	err = tc.taskModel.Update(task)
	if err != nil {
		data := Data{
			Info: err.Error(),
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusInternalServerError, response)
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

// DELETE /api/tasks/:id
func (tc taskController) DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := tc.taskModel.FindByID(id)
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

	err = tc.taskModel.Delete(task)
	if err != nil {
		data := Data{
			Info: err.Error(),
		}
		response := ResponseMessage{
			Status: "fail",
			Data:   data,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := ResponseMessage{
		Status: "success",
	}

	return c.JSON(http.StatusOK, response)
}
