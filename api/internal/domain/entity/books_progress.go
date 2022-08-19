package entity

import (
	"time"
)

type BooksProgress struct {
	ID        uint64 `gorm:"primary_key"`
	BooksId   uint64 `json:"books_id"`
	Page      uint64 `json:"page"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
