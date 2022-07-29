package main

import (
	"github.com/k1rnt/yonda/api/internal/interface/actions"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

// Init initializes the application
func Init() {
	// conn := setupDB()
	// v1 := e.Group("/api/v1")
	// {
	// 	v1.GET("/", actions.BookAllAction{Conn: conn}.Invoke)
	// 	v1.GET("/book/:id", actions.BookDetailAction{Conn: conn}.Invoke)
	// 	v1.POST("/book/:id", actions.BookReadAction{Conn: conn}.Invoke)
	// 	v1.POST("/book/delete/:id", actions.BookDeleteAction{Conn: conn}.Invoke)
	// 	v1.POST("/book/register", actions.BookRegisterAction{Conn: conn}.Invoke)
	// }
	e.GET("/ping", actions.PingAction{}.Invoke)
}

// Start starts the application
func Start() {
	e.Logger.Fatal(e.Start(":3000"))
}

// func setupDB() *gorm.DB {
// 	return config.Connect()
// }
