package main

import (
	"fmt"
	"time"
	// "io"
	// "net"
	// "time"
	// "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Countdown started.")
	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		moment := <-tick
		fmt.Println(moment.Format("15:04:05"), countdown)
	}
	fmt.Println("Rocket starts!")
}

//Возведение в квадрат - горутина посылает в канал целые числа,
//следующая горутина принимает из канала заначение и возводит ее в квадрат
//после чего результат направляет в канал, функция fmt выводит полученные
//из канала данные
// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	go func() {
// 		for i := 0; i <= 10; i++ {
// 			fmt.Print(i, "=",)
// 			naturals <- i
// 			time.Sleep(500 *time.Millisecond)
// 		}
// 		close(naturals)
// 	} ()

// 	go func() {
// 		for {
// 			n, ok := <- naturals
// 			if !ok {
// 				break
// 			}
// 			squares <- n * n
// 		}
// 		close(squares)
// 	} ()


// 	for {
// 		s, ok := <- squares
// 		if !ok {
// 			break
// 		}
// 		fmt.Println(s)
// 	}

// }




