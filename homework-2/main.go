package main

import (
	"fmt"
)

//Написать функцию, которая определяет, четное ли число.

func main() {
	fmt.Println("Вас приветствует программа проверки четности числа\nВведите число:")
	number := 0
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Println("Введенное число Четное")
	} else {
		fmt.Println("Введенное число Нечетное")
	}
	three()
}

//Написать функцию, которая определяет, делится ли число без остатка на 3.

func three() {
	fmt.Println("\nВас приветствует программа проверки деления на три без остатка\nВведите число:")
	number := 0
	fmt.Scanln(&number)
	if number%3 == 0 {
		fmt.Println("Введенное число делится на три без остатка")
	} else {
		fmt.Println("С этим числом проблемы")
	}
	fibonacci()
}

//Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.

func fibonacci() {
	fmt.Println("\nВас приветствует программа вывода на экран чисел Фибоначчи\nВведите колличество элементов в последовательности:")

	value := 0
	var fib int64 = 0
	var n1 int64 = 1
	var n2 int64 = 0

	fmt.Scanln(&value)
	fmt.Println(fib)

	for i := 0; i < value; i++ {
		n2 = n1
		n1 = fib
		fib = n1 + n2
		fmt.Println(fib)

	}
}
