package actions

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookDetailAction struct {
	Conn *gorm.DB
}

func (action BookDetailAction) Invoke(c echo.Context) error {
	id := c.Param("id")
	book := new(entity.Books)
	result := action.Conn.Where("id = ?", id).First(&book)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, &entity.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Book not found",
		})
	}
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, &book)
}
