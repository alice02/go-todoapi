package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/alice02/go_react_todoapp/controllers"
	"github.com/alice02/go_react_todoapp/database"
)

	
func main() {
	// initialize database
	database.InitDb()
	
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	api := e.Group("/api")
	api.GET("/todo", controllers.GetTodos)
	api.GET("/todo/:id", controllers.GetTodo)
	api.POST("/todo", controllers.PostTodo)
	api.PUT("/todo/:id", controllers.PutTodo)
	api.DELETE("/todo/:id", controllers.DeleteTodo)

	// listen on port 1323
	e.Logger.Debug(e.Start(":1323"))
}
