package main

import (
	"fmt"

	"./triangle"
)

func main() {
	var t triangle.Triangle
	//Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.

	t.SetCatet(14, 21)
	t.SetAltitude(10)
	fmt.Println(t.GetCatet())
	fmt.Println(t.Area())
}
