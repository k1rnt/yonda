package actions

import (
	"github.com/k1rnt/yonda/api/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PingAction struct{}

func (action PingAction) Invoke(c echo.Context) error {
	return c.JSON(http.StatusOK, &model.PingResponse{
		Status: "pong",
	})
}
