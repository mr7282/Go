package statistic

//Sum - this function returns the sum of all elements
func Sum(arrNum []float64) (float64) {
	var totalSum float64

	for _, num := range arrNum {
		totalSum += num
	}

	return totalSum
}