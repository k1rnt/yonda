package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
)

type BookRegisterDatastore struct {
	Conn *gorm.DB
}

func NewBookRegisterDatastore(conn *gorm.DB) BookRegisterDatastore {
	return BookRegisterDatastore{Conn: conn}
}

// Register is a function that registers the Book entity to the database.
func (db BookRegisterDatastore) Register(book *entity.Books) *gorm.DB {
	return db.Conn.Create(book)
}
