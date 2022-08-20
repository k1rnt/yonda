package query

import (
	"github.com/jinzhu/gorm"
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
)

type BookProgressRepository struct {
	Conn *gorm.DB
}

func NewBookProgressRepository(conn *gorm.DB) BookProgressRepository {
	return BookProgressRepository{Conn: conn}
}

func (conn BookProgressRepository) Run(id int) (*dto.BookProgress, *gorm.DB) {
	query := `
		select
			books_progresses.id,
			books_progresses.books_id,
			books_progresses.page
		from
		    books_progresses
		where
		    books_progresses.books_id = ?
	`
	var progress dto.BookProgress
	result := conn.Conn.Raw(query, id).Scan(&progress)
	return &progress, result
}

func (conn BookProgressRepository) Exist(id int) bool {
	_, result := conn.Run(id)
	return result.RowsAffected != 0
}
