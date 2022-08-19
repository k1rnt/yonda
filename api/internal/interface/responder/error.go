package responder

import "github.com/labstack/echo/v4"

type ErrorResponder struct {
	Status  int
	Message string
}

func NewErrorResponder(status int, message string) *ErrorResponder {
	return &ErrorResponder{
		Status:  status,
		Message: message,
	}
}

func (res ErrorResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
