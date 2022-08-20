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

type DeleteBookAction struct {
	Conn *gorm.DB
}

func (action DeleteBookAction) Invoke(c echo.Context) error {
	req := request.NewBookDeleteRequest()
	if err := req.Param(c, "id"); err != nil {
		return eresponder.NewErrorResponder(http.StatusBadRequest, err.Error()).Emit(c)
	}
	detail := usecase.NewBookDetailUsecase(action.Conn)
	if !detail.Exist(req.ID) {
		return eresponder.NewErrorResponder(http.StatusNotFound, "Book not found").Emit(c)
	}
	if err := usecase.NewDeleteBookUsecase(action.Conn).Delete(req.ID); err != nil {
		return eresponder.NewErrorResponder(http.StatusInternalServerError, "Failed to delete book").Emit(c)
	}
	return responder.NewDeleteBookResponder(http.StatusOK, "Book delete success").Emit(c)
}
