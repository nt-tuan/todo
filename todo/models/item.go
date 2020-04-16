package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Title  String `json:"title"`
	IsDone bool   `json:"is_done"`
}
