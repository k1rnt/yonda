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
	result := action.Conn.Where("deleted_at IS NULL").Find(&books)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, &entity.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Book not found",
		})
	}
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, &books)
}
