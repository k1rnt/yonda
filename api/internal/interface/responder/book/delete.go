package responder

import "github.com/labstack/echo/v4"

type DeleteBookResponder struct {
	Status  int
	Message string
}

func NewDeleteBookResponder(status int, message string) *DeleteBookResponder {
	return &DeleteBookResponder{
		Status:  status,
		Message: message,
	}
}

func (res *DeleteBookResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
