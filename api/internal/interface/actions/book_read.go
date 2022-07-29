package actions

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
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
	err := validation.Errors{
		"page": validation.Validate(bp.Page, validation.Required, validation.Min(1), validation.Max(book.PageCount)),
	}.Filter()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	existsProgress := new(entity.BooksProgress)
	pg := action.Conn.Where("books_id = ?", book.ID).First(&existsProgress)
	booksProgress := entity.BooksProgress{
		ID:      existsProgress.ID,
		BooksId: book.ID,
		Page:    bp.Page,
	}

	if pg.RowsAffected == 0 {
		// create book progress
		action.Conn.Create(&booksProgress)
	} else {
		// update book progress
		action.Conn.Model(&existsProgress).Where("books_id = ?", book.ID).Update("page", bp.Page)
	}
	return c.JSON(http.StatusOK, &book)
}
