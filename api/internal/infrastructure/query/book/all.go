package query

import (
	"github.com/jinzhu/gorm"
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
)

type AllBooksRepository struct {
	Conn *gorm.DB
}

func NewAllBooksRepository(conn *gorm.DB) AllBooksRepository {
	return AllBooksRepository{Conn: conn}
}

func (conn AllBooksRepository) Run() (*[]dto.BookDetail, error) {
	query := `
		select
			books.id,
			books.name,
			books.author,
			books.desc,
			books.page_count,
			books_progresses.page
		from
			books
		left join books_progresses on books.id = books_progresses.books_id
		where
			books.deleted_at is null
	`
	var books []dto.BookDetail
	err := conn.Conn.Raw(query).Scan(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}
