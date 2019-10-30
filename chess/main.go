package main

import (
	"fmt"
	"strconv"
)

//Range - преобразованный шахматный ряд
var Range []string

/////////////////////////////////////////// Почему не могу переместить "Position" и "Vp"
/////////////////////////////////////////// в функцию "GetPosition"

//Position - координаты позиции коня
type Position struct {
	X int
	Y int
}

//Vp - переменная позиции коня
var Vp Position

//GetPosition - ввод коорлинат первоначального нахождения коня
func GetPosition() {

	//S - переменная для преобразования
	var S string

	fmt.Println("Введите позицию КОНЯ по Оси Х (буквы от a до h)")
	fmt.Scanln(&S)
	Vp.X = convertInput(S)
	fmt.Println("Введите позицию КОНЯ по Оси Y")
	fmt.Scanln(&Vp.Y)
}

//converting - преобразование массива в шахматный ряд
func converting(k int) string {

	switch k {
	case 1:
		return "a"
	case 2:
		return "b"
	case 3:
		return "c"
	case 4:
		return "d"
	case 5:
		return "e"
	case 6:
		return "f"
	case 7:
		return "g"
	case 8:
		return "h"
	}
	return "-"
}

func convertInput(k string) int {
	switch k {
	case "a":
		return 1
	case "b":
		return 2
	case "c":
		return 3
	case "d":
		return 4
	case "e":
		return 5
	case "f":
		return 6
	case "g":
		return 7
	case "h":
		return 8
	}
	return 0 /////////////////////////////////Как избежать такой ситуации? (не нужный return)
}

func appendArray(l string) {
	Range = append(Range, l)
}

//ChessMove - расчет координат вариантов хода коня
func ChessMove() {

	type PositionMove struct {
		AxisX int
		AxisY int
	}

	ArrPM := []PositionMove{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}

	for _, arrrange := range ArrPM {

		newX := Vp.X + arrrange.AxisX
		newY := Vp.Y + arrrange.AxisY

		if newX > 0 && newX < 9 && newY > 0 && newY < 9 {

			appendArray(converting(newX) + strconv.Itoa(newY))
		} else {
			newX = 0
			appendArray(converting(newX))
		}
	}

}

func main() {

	GetPosition()
	ChessMove()
	fmt.Println(Range)

}
