package repository

import (
	"github.com/jinzhu/gorm"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
)

type BookRegisterRepository interface {
	Register(book *request.BookRegisterRequest) *gorm.DB
}
