package actions

import (
	"net/http"

	"github.com/jinzhu/gorm"
	eresponder "github.com/k1rnt/yonda/api/internal/interface/responder"
	responder "github.com/k1rnt/yonda/api/internal/interface/responder/book"
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
		return eresponder.NewErrorResponder(http.StatusInternalServerError, "Internal Server Error").Emit(c)
	}
	return responder.NewAllBookResponder(http.StatusOK, "Book all success", books).Emit(c)
}
