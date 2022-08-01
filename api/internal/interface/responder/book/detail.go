package responder

import (
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
	"github.com/labstack/echo/v4"
)

type DetailBookResponder struct {
	Status  int
	Message string
	Book    *[]dto.BookDetail
}

func NewDetailBookResponder(status int, message string, book *[]dto.BookDetail) *DetailBookResponder {
	return &DetailBookResponder{
		Status:  status,
		Message: message,
		Book:    book,
	}
}

func (res DetailBookResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
