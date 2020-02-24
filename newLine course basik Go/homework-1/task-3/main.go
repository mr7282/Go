package main

import (
	"fmt"
)

// Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.
//  *** Добавил возможность указывать длительность периода вклада самостоятельно.
func main() {
	fmt.Printf("Сумма влада через указанный период составит %v", getTotal(getSum(), getPercent(), getDuration()))
}

func getSum() int64 {
	fmt.Println("Введите желаемую сумму вклада:")
	var sum int64
	fmt.Scanln(&sum)
	return sum
}

func getPercent() int {
	fmt.Println("Введите желаемый процент:")
	var percent int
	fmt.Scanln(&percent)
	return percent
}

func getDuration() int {
	fmt.Println("Введите длительность периода вклада (в годах):")
	var duration int
	fmt.Scanln(&duration)
	return duration
}

func getTotal(sum int64, percent int, duration int) int64 {
	var total int64 = int64(sum)
	for i := 0; i < duration; i++ {
		total = total + ((total * int64(percent)) / 100)
	}
	return total
}
