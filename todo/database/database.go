package database

import (
	"github.com/thanhtuan260593/todo/models"
)

func (db *Database) CreateItem(title string) (*models.Item, error) {
	item := models.Item{
		Title:  title,
		IsDone: false,
	}
	if err := db.Create(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
