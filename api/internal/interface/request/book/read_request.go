package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
	"github.com/labstack/echo/v4"
)

type BookReadRequest struct {
	ID   string `json:"id"`
	Page int64  `json:"page"`
}

func NewBookReadRequest() *BookReadRequest {
	return &BookReadRequest{}
}

func (r *BookReadRequest) Param(c echo.Context, param string) error {
	r.ID = c.Param(param)
	if err := r.validateParam(); err != nil {
		return err
	}
	return nil
}

func (r *BookReadRequest) Bind(c echo.Context, book *dto.BookDetail) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err2 := r.validateBind(book); err2 != nil {
		return err2
	}
	return nil
}

func (r *BookReadRequest) validateParam() error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.ID,
			validation.Required.Error("ID is required"),
		),
	)
}

func (r *BookReadRequest) validateBind(book *dto.BookDetail) error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.Page,
			validation.Required.Error("Page is required"),
			validation.Min(1).Error("Page must be greater than 0"),
			validation.Max(book.PageCount).Error("Page must be less than book's page count"),
		),
	)
}
