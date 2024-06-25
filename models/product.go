package models

import (
	"strconv"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) SetValues(id int, name string, price float64) {
	p.ID = id
	p.Name = name
	p.Price = price
}

func (p *Product) ToString() string {
	return strconv.Itoa(p.ID) + " " + p.Name + " " + strconv.FormatFloat(p.Price, 'f', 2, 64)
}
