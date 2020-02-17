package main

import (
	"github.com/jeremylombogia/dewa/blog"
	"github.com/jeremylombogia/dewa/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", HelloWorld)

	e.GET("/user", user.User{}.Get)
	e.POST("/user", user.User{}.Post)

	e.GET("/blog", blog.Blog{}.Get)
	e.POST("/blog", blog.Blog{}.Post)

	e.Start(":1323")
}

// HelloWorld is index for Dewa
// to prevent denied access ping check
// for some cloud
func HelloWorld(c echo.Context) error {
	return c.String(200, "OK!")
}
