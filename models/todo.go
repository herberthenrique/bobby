package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
