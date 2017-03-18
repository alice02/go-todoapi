package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/alice02/go-todoapi/controllers"
	"github.com/alice02/go-todoapi/database"
)

func main() {
	// initialize database
	database.InitDb()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	api := e.Group("/api")
	api.GET("/tasks", controllers.GetTasks)
	api.GET("/tasks/:id", controllers.GetTask)
	api.POST("/tasks", controllers.PostTask)
	api.PUT("/tasks/:id", controllers.PutTask)
	api.DELETE("/tasks/:id", controllers.DeleteTask)

	// listen on port 1323
	e.Logger.Debug(e.Start(":1323"))
}
