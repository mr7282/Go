package main

import (
	"fmt"
)

func main() {

	rate := 64
	summ := 0
	howMuch := 0

	fmt.Println("Hi! Hom much you change money?")
	fmt.Scanln(&howMuch)
	summ = howMuch * rate
	fmt.Println("Сумма на руки- ", summ)

}
