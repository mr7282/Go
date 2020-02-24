package main

import (
	"fmt"
	"math"
)

// Даны катеты прямоугольного треугольника. Найти его площадь, периметр и
// гипотенузу. Используйте тип данных float64 и функции из пакета math.
func main() {
	var catetA float64 = 3
	var catetB float64 = 4
	fmt.Printf("Площадь треугольника равна %v\n", getArea(catetA, catetB))
	fmt.Printf("Периметр тругольника равен %v\n", getPerimetr(catetA, catetB))
	fmt.Printf("Гипотенуза треугольника равна %v\n", getHypotenuse(catetA, catetB))
}

func getArea(a float64, b float64) float64 {
	return (a * b) / 2
}

func getPerimetr(a float64, b float64) float64 {
	return a + b + getHypotenuse(a, b)
}

func getHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}
