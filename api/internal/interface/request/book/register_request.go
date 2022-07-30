package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type BookRegisterRequest struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Desc      string `json:"desc"`
	PageCount int    `json:"page_count"`
}

func NewBookRegisterRequest() *BookRegisterRequest {
	return &BookRegisterRequest{}
}

func (r *BookRegisterRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := r.validate(); err != nil {
		return err
	}
	return nil
}

func (r *BookRegisterRequest) ToBooks() *entity.Books {
	return &entity.Books{
		Name:      r.Name,
		Author:    r.Author,
		Desc:      r.Desc,
		PageCount: r.PageCount,
	}
}

func (r *BookRegisterRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 100).Error("Name must be between 1 and 100 characters"),
		),
		validation.Field(
			&r.Author,
			validation.Required.Error("Author is required"),
			validation.Length(1, 100).Error("Author must be between 1 and 100 characters"),
		),
		validation.Field(
			&r.Desc,
			validation.Length(1, 1000).Error("Desc must be between 1 and 1000 characters"),
		),
		validation.Field(
			&r.PageCount,
			validation.Required.Error("PageCount is required"),
			validation.Min(1).Error("PageCount must be greater than 0"),
		),
	)
}
