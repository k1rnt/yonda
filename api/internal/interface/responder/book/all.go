package responder

import (
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
	"github.com/labstack/echo/v4"
)

type BookAllResponder struct {
	Status  int
	Message string
	Books   *[]dto.BookDetail
}

func NewAllBookResponder(status int, message string, books *[]dto.BookDetail) *BookAllResponder {
	return &BookAllResponder{
		Status:  status,
		Message: message,
		Books:   books,
	}
}

func (res BookAllResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
