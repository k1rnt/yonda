package entity

type BookDetails struct {
	Name     string `json:"name"`
	Author   string `json:"author"`
	Desc     string `json:"desc"`
	Progress int    `json:"progress"`
}
