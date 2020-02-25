package main

import (
	"fmt"
)

// Написать функцию, которая определяет, четное ли число.
func main() {
	evenNumber(getNumber())
}

func getNumber() int {
	var number int
	fmt.Println("Введите число которое хотите проверить на четность")
	fmt.Scanln(&number)
	return number
}

func evenNumber(number int) {
	if (number % 2) == 0 {
		fmt.Printf("Введенное число четное")
	} else {
		fmt.Printf("Введенное число %v является не четным", number)
	}
}
