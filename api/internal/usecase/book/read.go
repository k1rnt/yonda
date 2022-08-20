package usecase

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
	"github.com/k1rnt/yonda/api/internal/domain/entity"
	datastore "github.com/k1rnt/yonda/api/internal/infrastructure/datastore/book"
	query "github.com/k1rnt/yonda/api/internal/infrastructure/query/book"
	request "github.com/k1rnt/yonda/api/internal/interface/request/book"
)

type BookReadUsecase struct {
	Conn *gorm.DB
}

func NewBookReadUsecase(conn *gorm.DB) BookReadUsecase {
	return BookReadUsecase{
		Conn: conn,
	}
}

func (u BookReadUsecase) Read(req *request.BookReadRequest) error {
	ds := datastore.NewReadBookDatastore(u.Conn)
	id, _ := strconv.ParseUint(req.ID, 10, 64)
	page := uint64(req.Page)

	var result *gorm.DB
	if !u.ProgressExist(id) {
		progress := entity.BooksProgress{
			BooksId:   id,
			Page:      page,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		result = ds.Create(progress)
	} else {
		result = ds.Update(&entity.BooksProgress{}, req)
	}
	return result.Error
}

func (u BookReadUsecase) ProgressExist(id uint64) bool {
	return query.NewBookProgressRepository(u.Conn).Exist(int(id))
}

func (u BookReadUsecase) BookExist(id uint64) bool {
	return query.NewBookDetailRepository(u.Conn).Exist(int(id))
}

func (u BookReadUsecase) GetBook(id uint64) *dto.BookDetail {
	q := query.NewBookDetailRepository(u.Conn)
	books, _ := q.Run(int(id))
	return books
}
