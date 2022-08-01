package actions

import (
	"github.com/jinzhu/gorm"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
	responder "github.com/k1rnt/yonda/api/internal/interface/responder/book"
	usecase "github.com/k1rnt/yonda/api/internal/usecase/book"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookDetailAction struct {
	Conn *gorm.DB
}

func (action BookDetailAction) Invoke(c echo.Context) error {
	req := request.NewBookDetailRequest()
	if err := req.Param(c, "id"); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	u := usecase.NewBookDetailUsecase(action.Conn)
	book, err := u.Detail(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(*book) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	resp := responder.NewDetailBookResponder(http.StatusOK, "Book detail success", book)
	return resp.Emit(c)
}
