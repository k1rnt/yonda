package request

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type BookDeleteRequest struct {
	ID int `json:"id"`
}

func NewBookDeleteRequest() *BookDeleteRequest {
	return &BookDeleteRequest{}
}

func (r *BookDeleteRequest) Param(c echo.Context, param string) error {
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

func (r *BookDeleteRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.ID,
			validation.Required.Error("ID is required"),
		),
	)
}
