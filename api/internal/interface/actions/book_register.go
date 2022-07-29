package actions

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
	responder "github.com/k1rnt/yonda/api/internal/interface/responder/book"
	usecase "github.com/k1rnt/yonda/api/internal/usecase/book"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// BookRegisterAction is action for book register
type BookRegisterAction struct {
	Conn *gorm.DB
}

// Invoke is action for book register
func (action BookRegisterAction) Invoke(c echo.Context) error {
	req := request.NewBookRegisterRequest()
	if err := req.Bind(c); err != nil {
		return err
	}
	books := req.ToBooks()
	u := usecase.NewRegisterBookUsecase(action.Conn)
	save := u.Register(books)
	if save.Error != nil {
		log.Println(save.Error)
		return c.JSON(http.StatusInternalServerError, &entity.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		})
	}
	resp := responder.NewRegisterBookResponder(http.StatusOK, "Book register success", books)
	return resp.Emit(c)
}
