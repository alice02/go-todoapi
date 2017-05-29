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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.Static("/", "build")

	// routing
	api := e.Group("/api")
	// tasks
	taskModel := models.NewTaskModel(db)
	taskController := controllers.NewTaskController(taskModel)
	api.GET("/tasks", taskController.GetTasks)
	api.GET("/tasks/:id", taskController.GetTask)
	api.POST("/tasks", taskController.PostTask)
	api.PUT("/tasks/:id", taskController.PutTask)
	api.DELETE("/tasks/:id", taskController.DeleteTask)
	// listen on port 1323
	e.Logger.Debug(e.Start(":1323"))
}
