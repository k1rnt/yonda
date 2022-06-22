package router

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/actions"
	"github.com/k1rnt/yonda/api/config"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func Init() {
	conn := setupDB()
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", actions.BookAllAction{}.Invoke)
		v1.POST("/book/register", actions.BookRegisterAction{Conn: conn}.Invoke)
	}
	e.GET("/ping", actions.PingAction{}.Invoke)
}

func Start() {
	e.Logger.Fatal(e.Start(":3000"))
}

func setupDB() *gorm.DB {
	return config.Connect()
}
