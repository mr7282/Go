package main

import (
	"fmt"
)

// Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
// Числа Фибоначчи определяются соотношениями Fn=Fn-1 + Fn-2.
func main() {
	fn := 0
	n1 := 1
	n2 := 0
	for {
		fmt.Println(fn)
		fn = n1 + n2
		n1 = n2
		n2 = fn
		if fn < 100 {
			continue
		} else {
			break
		}
	}
}
