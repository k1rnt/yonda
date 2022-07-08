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
	bookId := c.Param("id")
	var bp entity.BooksProgress
	if err := c.Bind(&bp); err != nil {
		return err
	}
	book := new(entity.Books)

	// book exists
	result := action.Conn.Where("id = ?", bookId).Where("deleted_at is null").First(&book)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, &entity.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Book not found",
		})
	}
	if result.Error != nil {
		return result.Error
	}

	existsProgress := new(entity.BooksProgress)
	pg := action.Conn.Where("books_id = ?", book.ID).First(&existsProgress)
	booksProgress := map[string]interface{}{
		"id":       existsProgress.ID,
		"books_id": book.ID,
		"progress": bp.Progress,
	}

	if pg.RowsAffected == 0 {
		// create book progress
		action.Conn.Create(&booksProgress)
	} else {
		// update book progress
		action.Conn.Model(&existsProgress).Where("books_id = ?", book.ID).Update("progress", bp.Progress)
	}
	return c.JSON(http.StatusOK, &book)
}
