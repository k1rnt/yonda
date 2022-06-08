package router

import (
	"github.com/k1rnt/yonda/api/actions"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func Init() {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", actions.BookAllAction{}.Invoke)
	}
	e.GET("/ping", actions.PingAction{}.Invoke)
}

func Start() {
	e.Logger.Fatal(e.Start(":3000"))
}
