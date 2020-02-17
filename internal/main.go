package internal

import "github.com/labstack/echo/v4"

// ControllerHandler contract
type ControllerHandler interface {
	Get(c echo.Context) error
	Post(c echo.Context) error
}

// RepositoryHandler contract
type RepositoryHandler interface {
	Select() interface{}
	Insert() interface{}
}

