package responder

import "github.com/labstack/echo/v4"

type PingResponder struct {
	Status  int
	Message string
}

func NewPingResponder(status int, message string) *PingResponder {
	return &PingResponder{
		Status:  status,
		Message: message,
	}
}

func (res PingResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
