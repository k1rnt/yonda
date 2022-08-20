package request

import (
	"strconv"

	"github.com/labstack/echo/v4"

	validation "github.com/go-ozzo/ozzo-validation"
)

type BookDetailRequest struct {
	ID int `json:"id"`
}

func NewBookDetailRequest() *BookDetailRequest {
	return &BookDetailRequest{}
}

func (r *BookDetailRequest) Param(c echo.Context, param string) error {
	var err error
	r.ID, err = strconv.Atoi(c.Param(param))
	if err != nil {
		return err
	}
	if err2 := r.validate(); err2 != nil {
		return err2
	}
	return nil
}

func (r *BookDetailRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.ID,
			validation.Required.Error("ID is required"),
		),
	)
}
