package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/alice02/go-todoapi/controllers"
	"github.com/alice02/go-todoapi/database"
	"github.com/alice02/go-todoapi/models"
)

func main() {
	// initialize database
	db, err := database.NewDB()
	if err != nil {
		panic("database connect failed")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	api := e.Group("/api")
	// tasks
	taskModel := models.NewTaskModel(db)
	taskController := controllers.NewTaskController(taskModel)
	api.GET("/tasks", taskController.GetTasks)
	/*
		api.GET("/tasks/:id", controllers.GetTask)
		api.POST("/tasks", controllers.PostTask)
		api.PUT("/tasks/:id", controllers.PutTask)
		api.DELETE("/tasks/:id", controllers.DeleteTask)
	*/
	// listen on port 1323
	e.Logger.Debug(e.Start(":1323"))
}
