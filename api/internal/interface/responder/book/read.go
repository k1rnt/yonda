package responder

import (
	"github.com/labstack/echo/v4"
)

type ReadBookResponder struct {
	Status  int
	Message string
}

func NewReadBookResponder(status int, message string) *ReadBookResponder {
	return &ReadBookResponder{
		Status:  status,
		Message: message,
	}
}

func (res ReadBookResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
