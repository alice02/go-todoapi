package main

import (
	"net/http"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
	

func todoList(c echo.Context) error {
	u := &User{
		Name: "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	g := e.Group("/api")
	g.GET("/", todoList)

	e.Logger.Debug(e.Start(":1323"))
}
