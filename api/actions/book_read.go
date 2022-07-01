package actions

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookReadAction struct {
	Conn *gorm.DB
}

func (action BookReadAction) Invoke(c echo.Context) error {
	id := c.Param("id")
	book := new(entity.Books)
	booksProgress := new(entity.BooksProgress)
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

	progress := action.Conn.Where("books_id = ?", book.ID).First(&booksProgress)
	if err := c.Bind(&booksProgress); err != nil {
		return err
	}
	if progress.RowsAffected == 0 {
		// create book progress
		action.Conn.Create(&booksProgress)
	} else {
		// update book progress
		action.Conn.Save(&booksProgress)
	}
	return c.JSON(http.StatusOK, &book)
}
