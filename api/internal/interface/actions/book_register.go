package actions

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type BookRegisterAction struct {
	Conn *gorm.DB
}

func (action BookRegisterAction) Invoke(c echo.Context) error {
	book := new(entity.Books)
	if err := c.Bind(&book); err != nil {
		return err
	}
	result := action.Conn.Create(&book)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return c.JSON(http.StatusOK, &result)
}
