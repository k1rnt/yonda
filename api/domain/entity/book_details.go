package entity

type BookDetails struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Desc      string `json:"desc"`
	PageCount int    `json:"page_count"`
	Page      int    `json:"page"`
}
