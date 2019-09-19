package main

import (
	"fmt"
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

//PositionOption - возвращаемый массив координат хода коня
var PositionOption []PositionMove

//Range - преобразованный шахматный ряд
var Range []string

//GetPosition - ввод коорлинат первоначального нахождения коня
func (p *Position) GetPosition() {
	fmt.Println("Введите позицию КОНЯ по Оси Х")
	fmt.Scanln(&p.x)
	fmt.Println("Введите позицию КОНЯ по Оси Y")
	fmt.Scanln(&p.y)
}

//ChessMove - расчет координат вариантов хода коня
func (pm PositionMove) ChessMove() {
	if vp.x <= 7 && vp.x >= 1 && vp.y <= 6 && vp.y >= 1 {
		vpm.pm1 = Position{x: vp.x + 1, y: vp.y + 2}

	} else {
		vpm.pm1 = Position{x: 0, y: 0}
	}

	if vp.x <= 6 && vp.x >= 1 && vp.y <= 7 && vp.y >= 1 {
		vpm.pm2 = Position{x: vp.x + 2, y: vp.y + 1}

	} else {
		vpm.pm2 = Position{x: 0, y: 0}
	}

	if vp.x <= 6 && vp.x >= 1 && vp.y <= 8 && vp.y >= 2 {
		vpm.pm3 = Position{x: vp.x + 2, y: vp.y - 1}

	} else {
		vpm.pm3 = Position{x: 0, y: 0}
	}

	if vp.x <= 7 && vp.x >= 1 && vp.y <= 8 && vp.y >= 3 {
		vpm.pm4 = Position{x: vp.x + 1, y: vp.y - 2}

	} else {
		vpm.pm4 = Position{x: 0, y: 0}
	}

	if vp.x <= 8 && vp.x >= 2 && vp.y <= 8 && vp.y >= 3 {
		vpm.pm5 = Position{x: vp.x - 1, y: vp.y - 2}

	} else {
		vpm.pm5 = Position{x: 0, y: 0}
	}

	if vp.x <= 8 && vp.x >= 3 && vp.y <= 8 && vp.y >= 2 {
		vpm.pm6 = Position{x: vp.x - 2, y: vp.y - 1}

	} else {
		vpm.pm6 = Position{x: 0, y: 0}
	}

	if vp.x <= 8 && vp.x >= 3 && vp.y <= 7 && vp.y >= 1 {
		vpm.pm7 = Position{x: vp.x - 2, y: vp.y + 1}

	} else {
		vpm.pm7 = Position{x: 0, y: 0}
	}

	if vp.x <= 8 && vp.x >= 2 && vp.y <= 6 && vp.y >= 1 {
		vpm.pm8 = Position{x: vp.x - 1, y: vp.y + 2}

	} else {
		vpm.pm8 = Position{x: 0, y: 0}
	}

}

//converting - преобразование массива в шахматный ряд
func converting() {

}

func main() {

	vp.GetPosition()
	vpm.ChessMove()
	PositionOption = append(PositionOption, vpm)
	fmt.Println(PositionOption)

}
