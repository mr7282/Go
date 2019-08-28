package main

import (
	"fmt"

	"./figures"
)

func main() {
	var t figures.Triangle
	var tz figures.Trapeze
	//Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.

	t.SetCatet()
	t.SetAltitude()
	t.GetCatet()
	fmt.Println("Площадь треугольника =", t.Area())
	tz.SetBasis()
	tz.SetAltitude()
	tz.GetBasis()
	fmt.Println
}
