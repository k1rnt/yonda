package usecase

import (
	"github.com/jinzhu/gorm"
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
	query "github.com/k1rnt/yonda/api/internal/infrastructure/query/book"
)

type AllBookUsecase struct {
	processor query.AllBooksRepository
}

func NewAllBookUsecase(conn *gorm.DB) AllBookUsecase {
	return AllBookUsecase{
		processor: query.NewAllBooksRepository(conn),
	}
}

func (u AllBookUsecase) All() (*[]dto.BookDetail, error) {
	books, err := u.processor.Run()
	if err != nil {
		return nil, err
	}
	return books, nil
}
