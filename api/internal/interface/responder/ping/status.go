package responder

import "github.com/labstack/echo/v4"

// PingResponder is a responder for ping
type PingResponder struct {
	Status  int
	Message string
}

// NewPingResponder creates a new PingResponder
func NewPingResponder(status int, message string) *PingResponder {
	return &PingResponder{
		Status:  status,
		Message: message,
	}
}

// Emit emits the response
func (res PingResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
