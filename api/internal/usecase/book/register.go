package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	datastore "github.com/k1rnt/yonda/api/internal/infrastructure/datastore/book"
)

type RegisterBookUsecase struct {
	Conn *gorm.DB
}

func NewRegisterBookUsecase(conn *gorm.DB) RegisterBookUsecase {
	return RegisterBookUsecase{Conn: conn}
}

func (u RegisterBookUsecase) Register(books *entity.Books) *gorm.DB {
	return datastore.NewBookRegisterDatastore(u.Conn).Register(books)
}
