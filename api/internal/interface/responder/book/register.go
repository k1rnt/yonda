package responder

import (
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type RegisterBookResponder struct {
	Status  int
	Message string
	Books   *entity.Books
}

func NewRegisterBookResponder(status int, message string, books *entity.Books) *RegisterBookResponder {
	return &RegisterBookResponder{
		Status:  status,
		Message: message,
		Books:   books,
	}
}

func (res RegisterBookResponder) Emit(c echo.Context) error {
	return c.JSON(res.Status, res)
}
