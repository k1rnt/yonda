package usecase

import (
	"github.com/jinzhu/gorm"
	datastore "github.com/k1rnt/yonda/api/internal/infrastructure/datastore/book"
)

type DeleteBookUsecase struct {
	Conn *gorm.DB
}

func NewDeleteBookUsecase(conn *gorm.DB) DeleteBookUsecase {
	return DeleteBookUsecase{Conn: conn}
}

func (u DeleteBookUsecase) Delete(id int) error {
	return datastore.NewDeleteBookDatastore(u.Conn).Delete(id)
}
