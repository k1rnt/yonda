package actions

import (
	"github.com/jinzhu/gorm"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
	responder "github.com/k1rnt/yonda/api/internal/interface/responder/book"
	usecase "github.com/k1rnt/yonda/api/internal/usecase/book"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BookReadAction struct {
	Conn *gorm.DB
}

func (action BookReadAction) Invoke(c echo.Context) error {
	req := request.NewBookReadRequest()
	if err := req.Param(c, "id"); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, _ := strconv.ParseUint(req.ID, 10, 64)
	u := usecase.NewBookReadUsecase(action.Conn)
	ok := u.BookExist(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}
	book := u.GetBook(id)
	if err := req.Bind(c, book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := u.Read(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := responder.NewReadBookResponder(http.StatusOK, "Book read success")
	return resp.Emit(c)
}
