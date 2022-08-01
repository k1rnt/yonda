package actions

import (
	responder "github.com/k1rnt/yonda/api/internal/interface/responder/ping"
	"github.com/labstack/echo/v4"
	"net/http"
)

// PingAction is action struct
type PingAction struct{}

// Invoke ping action
func (action PingAction) Invoke(c echo.Context) error {
	status := responder.NewPingResponder(http.StatusOK, "pong")
	return status.Emit(c)
}
