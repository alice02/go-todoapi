package controllers

import (
	"net/http"

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

/*
// GET /api/tasks/:id
func (ct controller) GetTask(c echo.Context) error {
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

// POST /api/tasks
func (ct controller) PostTask(c echo.Context) error {
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

// PUT /api/tasks/:id
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

// DELETE /api/tasks/:id
func DeleteTask(c echo.Context) error {
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
*/
