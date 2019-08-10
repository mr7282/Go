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
