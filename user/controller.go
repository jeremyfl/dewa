package user

import (
	"github.com/jeremylombogia/dewa/internal"
	"github.com/labstack/echo/v4"
)

type User struct {
	internal.ControllerHandler
	internal.RepositoryHandler
}

func (u User) Get(c echo.Context) error {
	return c.String(200, "OK")
}

func (u User) Post(c echo.Context) error {
	return c.String(200, "OK")
}
