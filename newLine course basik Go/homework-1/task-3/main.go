package main

import (
	"fmt"
)

// Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.
//  *** Добавил возможность указывать длительность периода вклада самостоятельно.
func main() {
	fmt.Printf("Сумма вклада через указанный период составит %v", getTotal(getSum(), getPercent(), getDuration()))
}

func getSum() float64 {
	fmt.Println("Введите желаемую сумму вклада:")
	var sum float64
	fmt.Scanln(&sum)
	return sum
}

func getPercent() float64 {
	fmt.Println("Введите желаемый процент:")
	var percent float64
	fmt.Scanln(&percent)
	return percent
}

func getDuration() int {
	fmt.Println("Введите длительность периода вклада (в годах):")
	var duration int
	fmt.Scanln(&duration)
	return duration
}

func getTotal(sum float64, percent float64, duration int) float64 {
	var total float64 = sum
	for i := 0; i < duration; i++ {
		total = total + ((total * percent) / 100)
	}
	return total
}
