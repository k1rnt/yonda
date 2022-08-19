package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
)

type ReadBookDatastore struct {
	Conn *gorm.DB
}

func NewReadBookDatastore(conn *gorm.DB) ReadBookDatastore {
	return ReadBookDatastore{Conn: conn}
}

func (db ReadBookDatastore) Create(progress entity.BooksProgress) *gorm.DB {
	return db.Conn.Create(&progress)
}

func (db ReadBookDatastore) Update(progress *entity.BooksProgress, req *request.BookReadRequest) *gorm.DB {
	return db.Conn.Model(progress).Where("books_id = ?", req.ID).Update("page", req.Page)
}
