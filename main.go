package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ControllerHandler
	RepositoryHandler
}

type Blog struct{
	ControllerHandler
	RepositoryHandler
}

type ControllerHandler interface {
	Get(c echo.Context) error
	Post(c echo.Context) error
}

type RepositoryHandler interface {
	Select() interface{}
	Insert() interface{}
}

func (u User) Get(c echo.Context) error {
	return c.String(200, "OK")
}

func (u User) Post(c echo.Context) error {
	return c.String(200, "OK")
}

func (b Blog) Get(c echo.Context) error {
	return c.String(200, "OK")
}

func (b Blog) Post(c echo.Context) error {
	return c.String(200, "OK")
}

func main() {
	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/user", User{}.Get)
	e.POST("/user", User{}.Post)

	e.GET("/blog", Blog{}.Get)
	e.POST("/blog", Blog{}.Post)

	e.GET("/", HelloWorld)

	e.Start(":1323")
}

func HelloWorld(c echo.Context) error {
	return c.String(200, "OK!")
}
