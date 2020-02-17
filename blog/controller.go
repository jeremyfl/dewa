package blog

import (
	"github.com/jeremylombogia/dewa/internal"
	"github.com/labstack/echo/v4"
)

type Blog struct {
	internal.ControllerHandler
	internal.RepositoryHandler
}

func (b Blog) Get(c echo.Context) error {
	return c.String(200, "OK")
}

func (b Blog) Post(c echo.Context) error {
	return c.String(200, "OK")
}
