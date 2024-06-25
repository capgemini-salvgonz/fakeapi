package models

import (
	"strconv"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (p *Product) SetValues(id int, name string, price float64) {
	p.ID = id
	p.Name = name
	p.Price = price
}

func (p *Product) ToString() string {
	return strconv.Itoa(p.ID) + " " + p.Name + " " + strconv.FormatFloat(p.Price, 'f', 2, 64)
}
