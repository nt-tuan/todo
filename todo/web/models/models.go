package model

import "github.com/thanhtuan260593/todo/core/entities"

//Pagable model
type Pagable struct {
	Page     uint   `form:"page"`
	PageSize uint   `form:"pageSize"`
	OrderBy  string `form:"orderBy"`
	OrderDir string `form:"orderDir"`
}

//ToEntity from page model
func (p *Pagable) ToEntity() *entities.Page {
	return &entities.Page{
		Page:     p.Page,
		PageSize: p.PageSize,
		Orders:   []entities.Order{*entities.NewOrder(p.OrderBy, p.OrderDir)},
	}
}
