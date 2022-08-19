package dto

type BookProgress struct {
	ID      int `json:"id" db:"id"`
	BooksID int `json:"books_id" db:"books_id"`
	Page    int `json:"page" db:"page"`
}
