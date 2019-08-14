package main

import (
	"fmt"
)

func main() {
	howMuch := 0
	rate := 64

	fmt.Println("Сколько хотите поменять $?")
	fmt.Scanln(&howMuch)
	howMuch = howMuch * rate
	fmt.Println("Сумма к выдаче- ", howMuch)

}

//Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу. Используйте тип данных float64 и функции из пакета math.

func triangle() {

	catetA := 0
	catet := 0

}
