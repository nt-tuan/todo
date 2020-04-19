package usecases

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/thanhtuan260593/todo/core/entities"
	infs "github.com/thanhtuan260593/todo/core/interfaces"
	"github.com/thanhtuan260593/todo/infrastructure/database"
)

var orderBys = []string{"title", "ID"}

//Todo ..
type Todo struct {
	db *database.Database
	database.Repository
}

//NewTodo create new todo
func NewTodo(db *database.Database) infs.ITodoUsecase {
	return &Todo{
		db: db,
	}
}

//GetValidOrderBy try to override getvalidorderby from repository
func (rep *Todo) GetValidOrderBy() []string {
	return []string{"id", "title", "is_done"}
}

//GetByID get item specified by its id
func (rep *Todo) GetByID(id uint) (*entities.Item, error) {
	var item entities.Item
	if err := rep.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

//ListItems get all items
func (rep *Todo) ListItems(page *entities.Page) ([]entities.Item, error) {
	var items []entities.Item
	var db *gorm.DB = rep.db.Table("items")
	db = rep.WherePage(db, page)
	if err := db.Find(&items).Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return items, nil
}

//CompleteItem ..
func (rep *Todo) CompleteItem(id uint) error {
	item, err := rep.GetByID(id)
	if err != nil {
		return err
	}

	item.IsDone = true
	if err := rep.db.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

//UpdateItem ..
func (rep *Todo) UpdateItem(id uint, p *entities.Item) (*entities.Item, error) {
	var item entities.Item
	if err := rep.db.First(&item, id).Error; err != nil {
		return &item, err
	}
	item.Title = p.Title
	if err := rep.db.Save(&item).Error; err != nil {
		return nil, err
	}
	return rep.GetByID(item.ID)
}

//ToggleItem ..
func (rep *Todo) ToggleItem(id uint) (*entities.Item, error) {
	var item entities.Item
	if err := rep.db.First(&item, id).Error; err != nil {
		return &item, err
	}
	item.IsDone = !item.IsDone
	if err := rep.db.Save(&item).Error; err != nil {
		return nil, err
	}

	if err := rep.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return rep.GetByID(item.ID)
}

//DeleteItem ...TodoRepos
func (rep *Todo) DeleteItem(id uint) error {
	var item entities.Item
	if err := rep.db.First(&item, id).Error; err != nil {
		return err
	}
	if err := rep.db.Delete(&item).Error; err != nil {
		return err
	}
	return nil
}

//CreateItem ...
func (rep *Todo) CreateItem(title string) (*entities.Item, error) {
	item := entities.Item{
		Title:  title,
		IsDone: false,
	}
	if err := rep.db.Create(&item).Error; err != nil {
		return nil, err
	}
	return rep.GetByID(item.ID)
}
