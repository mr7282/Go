package main

import (
	"fmt"
)

// !!!!!!!!!!!!Это 4 задание 4 урока. 3 задание пропущено по указанию препадователя!!!!!!!!!!!!!
//Написать функцию, которая будет получать позицию коня на шахматной доске, а возвращать массив из точек, в которые конь сможет сделать ход.
// a. Точку следует обозначить как структуру, содержащую x и y типа int
// b. Получение значений и их запись в точку должны происходить только с помощью отдельных методов. В них надо проводить проверку на то, что такая точка может существовать на шахматной доске.

type point struct {
	x int
	y int
}

type nextPoint point

func (p point) getPoint() point {
	getPoint := p
	positionHorse := ""
	fmt.Println("Введите позицию коня")
	fmt.Scanln(&positionHorse)
	getPoint.x, getPoint.y = convert(positionHorse)
	return getPoint
}

func convert(data string) (x, y int) {
	firstSymbol := string(data[0])
	secondSymbol := int(data[1])
	number := 0
	switch firstSymbol {
	case "a":
		return 

	}
	fmt.Println(firstSymbol)
	return firstSymbol, secondSymbol
}

func main() {
	myPoint := point{}
	fmt.Println(myPoint.getPoint())
}
