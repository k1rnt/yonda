package entity

import "github.com/jinzhu/gorm"

type BooksProgress struct {
	gorm.Model
	BookId   int `json:"book_id"`
	Progress int `json:"progress"`
}
