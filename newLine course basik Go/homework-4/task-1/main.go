package main

import (
	"fmt"
)

type product struct {
	id          int
	name        string
	price       int
	description string
}

type order struct {
	id         int
	content    []product
	totalPrice int
}

func (o *order) addProduct(data product) {
	o.content = append(o.content, data)
}

func (o *order) calculation() {
	for _, productPrice := range o.content {
		o.totalPrice += productPrice.price
	}
}

func main() {
	var myOrder order
	myOrder.id = 1
	myOrder.addProduct(product{1, "Холодильнк", 18000, "Холодильник Samsung"})
	myOrder.addProduct(product{2, "Стиральная машина", 12000, "Стиральная машина Wirpool"})
	myOrder.addProduct(product{3, "Чайник", 3000, "Чайник Tefal"})
	myOrder.calculation()
	fmt.Println(myOrder)
}
