package actions

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookAllAction struct {
	Conn *gorm.DB
}

func (action BookAllAction) Invoke(c echo.Context) error {
	var books []entity.Books
	action.Conn.Where("deleted_at IS NULL").Find(&books)
	return c.JSON(http.StatusOK, &books)
}
