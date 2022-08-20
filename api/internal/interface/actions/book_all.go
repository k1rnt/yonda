package actions

import (
	"net/http"

	"github.com/jinzhu/gorm"
	usecase "github.com/k1rnt/yonda/api/internal/usecase/book"
	"github.com/labstack/echo/v4"
)

type BookAllAction struct {
	Conn *gorm.DB
}

func (action BookAllAction) Invoke(c echo.Context) error {
	accessor := usecase.NewAllBookUsecase(action.Conn)
	books, err := accessor.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, books)
}
