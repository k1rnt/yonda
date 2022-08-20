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

// BookRegisterAction is action for book register
type BookRegisterAction struct {
	Conn *gorm.DB
}

// Invoke is action for book register
func (action BookRegisterAction) Invoke(c echo.Context) error {
	req := request.NewBookRegisterRequest()
	if err := req.Bind(c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	books := req.ToBooks()
	u := usecase.NewRegisterBookUsecase(action.Conn)
	save := u.Register(books)
	if save.Error != nil {
		return eresponder.NewErrorResponder(http.StatusInternalServerError, "Internal Server Error").Emit(c)
	}
	resp := responder.NewRegisterBookResponder(http.StatusOK, "Book register success", books)
	return resp.Emit(c)
}
