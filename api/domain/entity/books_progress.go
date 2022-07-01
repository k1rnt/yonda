package entity

import (
	"time"
)

type BooksProgress struct {
	ID        uint `gorm:"primary_key"`
	BooksId   uint `json:"books_id"`
	Progress  int  `json:"progress"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
