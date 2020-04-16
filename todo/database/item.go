package database

import "github.com/thanhtuan260593/todo/models"

//ListAllItems get all items
func (db *Database) ListAllItems() ([]models.Item, error) {
	var items []models.Item
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

//CompleteItem ..
func (db *Database) CompleteItem(id uint) error {
	var item models.Item
	if err := db.First(&item, id).Error; err != nil {
		return err
	}
	item.IsDone = true
	if err := db.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

//ToggleItem ..
func (db *Database) ToggleItem(id uint) error {
	var item models.Item
	if err := db.First(&item, id).Error; err != nil {
		return err
	}
	item.IsDone = !item.IsDone
	if err := db.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

//DeleteItem ...
func (db *Database) DeleteItem(id uint) error {
	var item models.Item
	if err := db.First(&item, id).Error; err != nil {
		return err
	}
	if err := db.Delete(&item).Error; err != nil {
		return err
	}
	return nil
}

//CreateItem ...
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
