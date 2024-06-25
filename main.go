package main

import (
	"github.com/chava.gnolasco/fakeapi/models"

	"log"
)

var products = make([]models.Product, 0)

func createProducts(id int, name string, price float64) {
	product := new(models.Product)
	product.SetValues(id, name, price)
	products = append(products, *product)
}

func main() {

	createProducts(1, "Product 1", 100.00)
	createProducts(2, "Product 2", 200.00)
	createProducts(3, "Product 3", 300.00)

	for index, product := range products {
		log.Printf(" - Product [%d] : [%s]\n", index, product.ToString())
	}
}
