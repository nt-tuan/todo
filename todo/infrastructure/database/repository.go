package database

import (
	"github.com/jinzhu/gorm"
	"github.com/thanhtuan260593/todo/core/entities"
)

//Repository define abstract query
type Repository struct {
}

var orderBys = []string{"id"}

//GetValidOrderBy helper
func (rep *Repository) GetValidOrderBy() []string {
	return orderBys
}

//IsValidOrderBy helper
func (rep *Repository) IsValidOrderBy(or string) bool {
	for _, b := range rep.GetValidOrderBy() {
		if b == or {
			return true
		}
	}
	return false
}

//WhereOrder query part
func (rep *Repository) WhereOrder(db *gorm.DB, ord *entities.Order) *gorm.DB {
	if rep.IsValidOrderBy(ord.By) {
		db = db.Order(ord.By + " " + ord.Dir)
	}
	return db
}

//WherePage query part
func (rep *Repository) WherePage(db *gorm.DB, page *entities.Page) *gorm.DB {
	if page.Orders != nil {
		for _, or := range page.Orders {
			db = rep.WhereOrder(db, &or)
		}
	}
	db = db.Offset(page.Page * page.PageSize).Limit(page.PageSize)
	return db
}

//WhereNext query part
func (rep *Repository) WhereNext(db *gorm.DB, p *entities.Next) *gorm.DB {
	db = db.Where(p.Condition)
	if p.Ignore != nil {
		db = db.Not("ID", p.Ignore)
	}
	db = db.Take(p.Length)
	return db
}
