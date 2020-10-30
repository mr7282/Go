package main

import (
	"fmt"
	// "time"
)

func main() {
	// go spinner(50 * time.Millisecond)
	// time.Sleep(10 * time.Second) // Метод выполнения первого задания
	const n uint64 = 20
	// fibN := fibonacci(n)
	fibonacci(n)
	// fmt.Printf("\rFibonacci %d = %d\n", n, fibN)
}

// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range "-\\|/" {
// 			fmt.Printf("%c\r", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }

func fibonacci(num uint64) { //Данная реализация вычесления числового ряда Фибоначчии работает в разы быстрее
	// if num < 2 {            //нежили реализация через рекурсивный метод вызов функций
	// 	return num
	// }
	// return fibonacci(num-1) + fibonacci(num-2)

	var f []uint64
	var i uint64
	for i = 0; i <= num; i++ {
		if i < 2 {
			f = append(f, i)
		} else {
			f = append(f, (f[len(f)-2] + f[len(f)-1]))
		}
	}
	fmt.Println(f[len(f)-1])
}
