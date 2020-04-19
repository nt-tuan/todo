package infs

import (
	"github.com/thanhtuan260593/todo/core/entities"
)

//ITodoUsecase todo interfaces.
type ITodoUsecase interface {
	GetByID(uint) (*entities.Item, error)
	ListItems(*entities.Page) ([]entities.Item, error)
	CompleteItem(uint) error
	ToggleItem(uint) (*entities.Item, error)
	UpdateItem(id uint, e *entities.Item) (*entities.Item, error)
	DeleteItem(uint) error
	CreateItem(string) (*entities.Item, error)
}
