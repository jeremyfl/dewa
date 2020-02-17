package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerHandler interface {
	Get(c echo.Context) error
	Post(c echo.Context) error
}

type UserController struct {
	Name string
}

func (u UserController) Get(c echo.Context) error {
	return c.String(200, "OK")
}

func (u UserController) Post(c echo.Context) error {
	return c.String(200, "OK")
}

func main() {
	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var controller ControllerHandler

	controller = UserController{"User Controller"}
	e.GET("/users", controller.Get)
	e.POST("/users", controller.Post)

	e.Start(":1323")
}
