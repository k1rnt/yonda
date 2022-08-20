package actions

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
	eresponder "github.com/k1rnt/yonda/api/internal/interface/responder"
	responder "github.com/k1rnt/yonda/api/internal/interface/responder/book"
	usecase "github.com/k1rnt/yonda/api/internal/usecase/book"
	"github.com/labstack/echo/v4"
)

type BookReadAction struct {
	Conn *gorm.DB
}

func (action BookReadAction) Invoke(c echo.Context) error {
	req := request.NewBookReadRequest()
	if err := req.Param(c, "id"); err != nil {
		resp := eresponder.NewErrorResponder(http.StatusBadRequest, err.Error())
		return resp.Emit(c)
	}
	id, _ := strconv.ParseUint(req.ID, 10, 64)
	u := usecase.NewBookReadUsecase(action.Conn)
	ok := u.BookExist(id)
	if !ok {
		resp := eresponder.NewErrorResponder(http.StatusNotFound, "Book not found")
		return resp.Emit(c)
	}
	book := u.GetBook(id)
	if err := req.Bind(c, book); err != nil {
		resp := eresponder.NewErrorResponder(http.StatusBadRequest, "page must be integer")
		return resp.Emit(c)
	}
	if err := u.Read(req); err != nil {
		resp := eresponder.NewErrorResponder(http.StatusInternalServerError, err.Error())
		return resp.Emit(c)
	}

	resp := responder.NewReadBookResponder(http.StatusOK, "Book read success")
	return resp.Emit(c)
}
