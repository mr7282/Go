package main

import (
	"fmt"
)

// 1. Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, //..открыты ли окна, насколько заполнен объем багажника.
// 2. Инициализировать несколько экземпляров структур. Применить к ним различные действия. Вывести значения свойств экземпляров в консоль.

type car struct {
	brand  string
	model  string
	color  string
	year   int
	engine string
}

type printer struct {
	brand       string
	model       string
	printerType string
	printType   string
	price       int
}

func main() {
	myCar := car{
		brand:  "Toyota",
		model:  "Siena",
		color:  "Silver",
		year:   2016,
		engine: "engine started",
	}

	myPrinter := printer{
		brand:       "Brother",
		model:       "DCP-7070DWR",
		printerType: "Multifunction",
		printType:   "Laser",
		price:       123,
	}

	fmt.Printf("Марка: %v,\nМодель: %v,\nЦвет: %v,\nГод выпуска: %v,\nДвигатель: %v\n\n", myCar.brand, myCar.model, myCar.color, myCar.year, myCar.engine)
	fmt.Printf("Производитель: %v\nМодель: %v\nТип устройства: %v\nТип печати: %v\nЦена: %v$", myPrinter.brand, myPrinter.model, myPrinter.printerType, myPrinter.printType, myPrinter.price)
}
