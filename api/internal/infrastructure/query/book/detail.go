package query

import (
	"github.com/jinzhu/gorm"
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
)

type BookDetailRepository struct {
	Conn *gorm.DB
}

func NewBookDetailRepository(conn *gorm.DB) BookDetailRepository {
	return BookDetailRepository{Conn: conn}
}

func (conn BookDetailRepository) Run(id int) (*dto.BookDetail, *gorm.DB) {
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
		    books.id = ?
		and
		    books.deleted_at is null
	`
	var book dto.BookDetail
	result := conn.Conn.Raw(query, id).Scan(&book)
	return &book, result
}

func (conn BookDetailRepository) Exist(id int) bool {
	_, result := conn.Run(id)
	return result.RowsAffected != 0
}
