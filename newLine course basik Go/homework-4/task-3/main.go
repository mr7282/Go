package main

import (
	"fmt"
)

// !!!!!!!!!!!!Это 4 задание 4 урока. 3 задание пропущено по указанию препадователя!!!!!!!!!!!!!
//Написать функцию, которая будет получать позицию коня на шахматной доске, а возвращать массив из точек, в которые конь сможет сделать ход.
// a. Точку следует обозначить как структуру, содержащую x и y типа int
// b. Получение значений и их запись в точку должны происходить только с помощью отдельных методов. В них надо проводить проверку на то, что такая точка может существовать на шахматной доске.

//point - координаты точки фигуры
type point struct {
	x int
	y int
}

//getPoint - получает значение точки в виде строки, конвертирует интерпритацию байт кода ASCII в числовые координаты,
//проверяет на возможность существования на шахматной доске, возвращает эти числовые координаты точки
func (p point) getPoint() point {
	userPoint := p
	positionHorse := ""

	for {
		fmt.Println("Введите позицию коня")
		fmt.Scanln(&positionHorse)
		userPoint.x, userPoint.y = int(positionHorse[0]), int(positionHorse[1])
		if userPoint.y > 48 && userPoint.y < 57 && userPoint.x > 96 && userPoint.x < 105 {
			break
		} else {
			fmt.Printf("Введена неправильная позиция. Попробуй еще раз!\n\n")
		}
	}
	return userPoint
}

//allowedMoveHorse - проверяет все варианты хода коня на допустимость и выводит в терминал точки не выходящие
//за пределы шахматной доски
func (p point) allowedMoveHorse() {
	move := []point{
		{p.x + 1, p.y + 2},
		{p.x + 2, p.y + 1},
		{p.x + 2, p.y - 1},
		{p.x + 1, p.y - 2},
		{p.x - 1, p.y - 2},
		{p.x - 2, p.y - 1},
		{p.x - 2, p.y + 1},
		{p.x - 1, p.y + 2},
	}

	for _, show := range move {
		if show.y > 48 && show.y < 57 && show.x > 96 && show.x < 105 {
			fmt.Printf("%v%v\n", string(show.x), string(show.y))
		}
	}
}

func main() {
	myPoint := point{}
	myPoint = myPoint.getPoint()
	myPoint.allowedMoveHorse()

}
