package entities

import "strings"

//Order param
type Order struct {
	By  string
	Dir string
}

//NewOrder init
func NewOrder(by string, dir string) *Order {
	order := Order{"", "asc"}
	if strings.ToLower(dir) == "desc" {
		order.Dir = "desc"
	}
	return &order
}

//Page entity param
type Page struct {
	Page     uint
	PageSize uint
	Orders   []Order
}

//Next entity param
type Next struct {
	Length    uint
	Condition uint
	Ignore    []uint
}
