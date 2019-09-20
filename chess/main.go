package main

import (
	"fmt"
	"strconv"
)

//Position - координаты позиции коня
type Position struct {
	x int
	y int
}

//P - переменная позиции коня
var vp Position

//PositionMove - координаты вариантов хода коня
type PositionMove struct {
	pm1 Position
	pm2 Position
	pm3 Position
	pm4 Position
	pm5 Position
	pm6 Position
	pm7 Position
	pm8 Position
}

//PM - переменная структуры вариантов хода
var vpm PositionMove

//Range - преобразованный шахматный ряд
var Range []string

//X - переменная для преобразования
var X string

//S - переменная для итоговой строки
var S string

//I - переменная инт
var I int

//GetPosition - ввод коорлинат первоначального нахождения коня
func (p *Position) GetPosition() {
	fmt.Println("Введите позицию КОНЯ по Оси Х (буквы от a до h)")
	fmt.Scanln(&S)
	convertInput(S)
	p.x = I
	fmt.Println("Введите позицию КОНЯ по Оси Y")
	fmt.Scanln(&p.y)
}

//ChessMove - расчет координат вариантов хода коня
func (pm PositionMove) ChessMove() {
	if vp.x <= 7 && vp.x >= 1 && vp.y <= 6 && vp.y >= 1 {
		vpm.pm1 = Position{x: vp.x + 1, y: vp.y + 2}
		converting(vpm.pm1.x)
		X = S + strconv.Itoa(vpm.pm1.y)
		appendArray(X)

	} else {
		vpm.pm1 = Position{x: 0, y: 0}
		converting(vpm.pm1.x)
		appendArray(S)
	}

	if vp.x <= 6 && vp.x >= 1 && vp.y <= 7 && vp.y >= 1 {
		vpm.pm2 = Position{x: vp.x + 2, y: vp.y + 1}
		converting(vpm.pm2.x)
		X = S + strconv.Itoa(vpm.pm2.y)
		appendArray(X)

	} else {
		vpm.pm2 = Position{x: 0, y: 0}
		converting(vpm.pm2.x)
		appendArray(S)
	}

	if vp.x <= 6 && vp.x >= 1 && vp.y <= 8 && vp.y >= 2 {
		vpm.pm3 = Position{x: vp.x + 2, y: vp.y - 1}
		converting(vpm.pm3.x)
		X = S + strconv.Itoa(vpm.pm3.y)
		appendArray(X)

	} else {
		vpm.pm3 = Position{x: 0, y: 0}
		converting(vpm.pm3.x)
		appendArray(S)
	}

	if vp.x <= 7 && vp.x >= 1 && vp.y <= 8 && vp.y >= 3 {
		vpm.pm4 = Position{x: vp.x + 1, y: vp.y - 2}
		converting(vpm.pm4.x)
		X = S + strconv.Itoa(vpm.pm4.y)
		appendArray(X)

	} else {
		vpm.pm4 = Position{x: 0, y: 0}
		converting(vpm.pm4.x)
		appendArray(S)
	}

	if vp.x <= 8 && vp.x >= 2 && vp.y <= 8 && vp.y >= 3 {
		vpm.pm5 = Position{x: vp.x - 1, y: vp.y - 2}
		converting(vpm.pm5.x)
		X = S + strconv.Itoa(vpm.pm5.y)
		appendArray(X)

	} else {
		vpm.pm5 = Position{x: 0, y: 0}
		converting(vpm.pm5.x)
		appendArray(S)
	}

	if vp.x <= 8 && vp.x >= 3 && vp.y <= 8 && vp.y >= 2 {
		vpm.pm6 = Position{x: vp.x - 2, y: vp.y - 1}
		converting(vpm.pm6.x)
		X = S + strconv.Itoa(vpm.pm6.y)
		appendArray(X)

	} else {
		vpm.pm6 = Position{x: 0, y: 0}
		converting(vpm.pm6.x)
		appendArray(S)
	}

	if vp.x <= 8 && vp.x >= 3 && vp.y <= 7 && vp.y >= 1 {
		vpm.pm7 = Position{x: vp.x - 2, y: vp.y + 1}
		converting(vpm.pm7.x)
		X = S + strconv.Itoa(vpm.pm7.y)
		appendArray(X)

	} else {
		vpm.pm7 = Position{x: 0, y: 0}
		converting(vpm.pm7.x)
		appendArray(S)
	}

	if vp.x <= 8 && vp.x >= 2 && vp.y <= 6 && vp.y >= 1 {
		vpm.pm8 = Position{x: vp.x - 1, y: vp.y + 2}
		converting(vpm.pm8.x)
		X = S + strconv.Itoa(vpm.pm8.y)
		appendArray(X)

	} else {
		vpm.pm8 = Position{x: 0, y: 0}
		converting(vpm.pm8.x)
		appendArray(S)
	}

}

//converting - преобразование массива в шахматный ряд
func converting(k int) {

	switch k {
	case 0:
		S = "-"
	case 1:
		S = "a"
	case 2:
		S = "b"
	case 3:
		S = "c"
	case 4:
		S = "d"
	case 5:
		S = "e"
	case 6:
		S = "f"
	case 7:
		S = "g"
	case 8:
		S = "h"
	}

}

func convertInput(k string) {
	switch k {
	case "a":
		I = 1
	case "b":
		I = 2
	case "c":
		I = 3
	case "d":
		I = 4
	case "e":
		I = 5
	case "f":
		I = 6
	case "g":
		I = 7
	case "h":
		I = 8
	}
}

func appendArray(l string) {
	Range = append(Range, l)
}

func main() {

	vp.GetPosition()
	vpm.ChessMove()
	fmt.Println(Range)

}
