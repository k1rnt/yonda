package dto

type BookDetail struct {
	Name      string `json:"name" db:"name"`
	Author    string `json:"author" db:"author"`
	Desc      string `json:"desc" db:"desc"`
	PageCount int    `json:"page_count" db:"page_count"`
	Page      int    `json:"page" db:"page"`
}
