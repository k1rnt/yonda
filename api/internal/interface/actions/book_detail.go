package actions

import (
	"net/http"

	"github.com/jinzhu/gorm"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
	eresponder "github.com/k1rnt/yonda/api/internal/interface/responder"
	responder "github.com/k1rnt/yonda/api/internal/interface/responder/book"
	usecase "github.com/k1rnt/yonda/api/internal/usecase/book"
	"github.com/labstack/echo/v4"
)

type BookDetailAction struct {
	Conn *gorm.DB
}

func (action BookDetailAction) Invoke(c echo.Context) error {
	req := request.NewBookDetailRequest()
	if err := req.Param(c, "id"); err != nil {
		return eresponder.NewErrorResponder(http.StatusBadRequest, err.Error()).Emit(c)
	}
	u := usecase.NewBookDetailUsecase(action.Conn)
	book, _ := u.Detail(req.ID)
	if !u.Exist(req.ID) {
		return eresponder.NewErrorResponder(http.StatusNotFound, "Book not found").Emit(c)
	}
	resp := responder.NewDetailBookResponder(http.StatusOK, "Book detail success", book)
	return resp.Emit(c)
}
