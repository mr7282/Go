package main

import (
	"fmt"
)

//Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.

type autopart struct {
	id          int
	name        string
	price       int
	description string
}

type invoiceAutopart struct {
	id      int
	content []autopart
}

type work struct {
	id       int
	name     string
	brand    string
	price    int
	normHour float32
}

type invoiceWork struct {
	id      int
	content []work
}

type total interface {
	Calculation() int
}

func (i *invoiceWork) addWork(data work) {
	i.content = append(i.content, data)
}

func (i *invoiceWork) Calculation() int {
	totalPrice := 0
	for _, workPrice := range i.content {
		totalPrice += workPrice.price
	}
	return totalPrice
}

func (o *invoiceAutopart) addAutoPart(data autopart) {
	o.content = append(o.content, data)
}

func (o *invoiceAutopart) Calculation() int {
	totalPrice := 0
	for _, AutoPartPrice := range o.content {
		totalPrice += AutoPartPrice.price
	}
	return totalPrice
}

func fullPrice(calcs ...total) int {
	res := 0
	for _, calc := range calcs {
		res += calc.Calculation()
	}
	return res
}

func main() {
	var myInvoiceAutopart invoiceAutopart
	myInvoiceAutopart.id = 1
	myInvoiceAutopart.addAutoPart(autopart{1, "Водяная помпа", 18000, "Audi, Skoda, VW"})
	myInvoiceAutopart.addAutoPart(autopart{2, "Распредвал TRW", 12000, "Audi, Skoda, VW"})
	myInvoiceAutopart.addAutoPart(autopart{3, "Щетки стеклоочистителя Bosch 60\"", 3000, "Multibrand"})
	var myInvoiceWork invoiceWork
	myInvoiceWork.id = 1
	myInvoiceWork.addWork(work{1, "Замена водяной помпы", "Skoda", 6000, 3.5})
	myInvoiceWork.addWork(work{2, "Замена распредвала", "Skoda", 8000, 4.5})
	myInvoiceWork.addWork(work{3, "Замена щеток стеклоочестителч", "Skoda", 1000, 0.5})
	fmt.Println(myInvoiceAutopart, "\n", myInvoiceAutopart.Calculation(), "\n", myInvoiceWork, "\n", myInvoiceWork.Calculation(), "\n", fullPrice(&myInvoiceWork, &myInvoiceAutopart))
}
