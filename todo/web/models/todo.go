package model

import "github.com/thanhtuan260593/todo/core/entities"

//ActionItem model binding from param
type ActionItem struct {
	ID uint `uri:"id" binding:"required"`
}

//ResponseItem response model
type ResponseItem struct {
	ID     uint
	Title  string
	IsDone bool
}

//NewResponseItem new
func NewResponseItem(e *entities.Item) *ResponseItem {
	return &ResponseItem{
		ID:     e.ID,
		Title:  e.Title,
		IsDone: e.IsDone,
	}
}

//CreateItem model
type CreateItem struct {
	Title string `json:"title"`
}
