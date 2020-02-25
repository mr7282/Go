package main

import (
	"fmt"
)

// Написать функцию, которая определяет, делится ли число без остатка на 3.
func main() {
	divisionOnThree(getNumber())
}

func getNumber() int {
	var number int
	fmt.Println("Введите число для проверки деления на 3")
	fmt.Scanln(&number)
	return number
}

func divisionOnThree(number int) {
	if (number % 3) == 0 {
		fmt.Printf("Число %v делится на 3 без остатка", number)
	} else {
		fmt.Printf("Число %v не делится на 3 без остатка", number)
	}
}
