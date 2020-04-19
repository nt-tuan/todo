package entities

import (
	"github.com/jinzhu/gorm"
)

//Item of to-do list
type Item struct {
	gorm.Model
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
