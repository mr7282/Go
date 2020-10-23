package main

import (
	"fmt"
	"task-1/statistic"
)

func main() {
	arr := []float64{1, 2, 3, 5, 7, 5}
	fmt.Println(statistic.Average(arr))
	fmt.Println(statistic.Sum(arr))
	fmt.Println(statistic.Len(arr))
}
