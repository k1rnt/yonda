package actions

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookDetailAction struct {
	Conn *gorm.DB
}

func (action BookDetailAction) Invoke(c echo.Context) error {
	id := c.Param("id")
	book := new(entity.BookDetails)
	result := action.Conn.Table("books").Select(
		[]string{"books.id", "books.name", "books.author", "books.desc", "books.page_count", "books_progresses.page"}).Joins(
		"LEFT JOIN books_progresses ON books.id = books_progresses.books_id").Where("books.id = ?", id).Where("books.deleted_at IS NULL").Scan(&book)
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
