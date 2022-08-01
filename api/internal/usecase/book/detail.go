package usecase

import (
	"github.com/jinzhu/gorm"
	dto "github.com/k1rnt/yonda/api/internal/domain/dto/book"
	query "github.com/k1rnt/yonda/api/internal/infrastructure/query/book"
)

type BookDetailUsecase struct {
	processor query.BookDetailRepository
}

func NewBookDetailUsecase(conn *gorm.DB) BookDetailUsecase {
	return BookDetailUsecase{
		processor: query.NewBookDetailRepository(conn),
	}
}

func (u BookDetailUsecase) Detail(id int) (*[]dto.BookDetail, error) {
	book, err := u.processor.Run(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
