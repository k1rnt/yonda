package entity

import "github.com/jinzhu/gorm"

type Books struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Desc      string `json:"desc"`
	PageCount int    `json:"page_count"`
}
