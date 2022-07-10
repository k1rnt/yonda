package actions

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type BookDeleteAction struct {
	Conn *gorm.DB
}

func (action BookDeleteAction) Invoke(c echo.Context) error {
	id := c.Param("id")
	book := new(entity.Books)
	result := action.Conn.Where("id = ?", id).Where("books.deleted_at IS NULL").First(&book)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, &entity.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Book not found",
		})
	}
	if result.Error != nil {
		return result.Error
	}
	err := action.Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&book).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
			return err
		}
		progress := new(entity.BooksProgress)
		action.Conn.Where("books_id = ?", id).First(&progress)
		if err := tx.Model(&entity.BooksProgress{}).Where("books_id = ?", id).Delete(&progress).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &book)
}
