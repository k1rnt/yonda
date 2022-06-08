package actions

import "github.com/labstack/echo/v4"

type BookAllAction struct{}

func (action BookAllAction) Invoke(c echo.Context) error {
	return nil
}
