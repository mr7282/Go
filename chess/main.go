package main

import (
	"fmt"
)

type position struct {
	x int
	y int
}

type PositionMove struct {
	x int
	y int
}

var PositionOption []PositionMove

func (p position) GetPosition() {
	fmt.Println("Введите позицию КОНЯ по Оси Х")
	fmt.Scanln(&p.x)
	fmt.Println("Введите позицию КОНЯ по Оси Y")
	fmt.Scanln(&p.y)
}

func (pm PositionMove) ChessMove() {

}

var x = position{x: 2, y: 3}

func main() {

	fmt.Println(x)
}
