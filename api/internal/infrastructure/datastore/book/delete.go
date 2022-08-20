package datastore

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
)

type DeleteBookDatastore struct {
	Conn *gorm.DB
}

func NewDeleteBookDatastore(conn *gorm.DB) DeleteBookDatastore {
	return DeleteBookDatastore{Conn: conn}
}

func (db DeleteBookDatastore) Delete(id int) error {
	err := db.Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entity.BooksProgress{}, "books_id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Model(&entity.Books{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
