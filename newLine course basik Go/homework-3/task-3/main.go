package main

import (
	"fmt"
	"queue" // Этот пакет создан в учебных целях. Он рапологается в GOPATH..../src/
)

// Реализовать очередь. Это структура данных, работающая по принципу FIFO (First Input — first output, или «первым зашел — первым вышел»).

// Выполнение задания заключалось в создании пакета очереди по принципу FIFO.
// Эта функция показывает лишь конечный результат.

func main() {
	queue.Push("1")
	queue.Push("2")
	queue.Push("3")
	queue.Push("4")
	queue.Push("5")
	queue.Push("6")
	queue.Push("7")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
