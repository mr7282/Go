package statistic

//Average - this function returns the average of the array
func Average(arrNum []float64) float64 {
	total := float64(0)

	for _, num := range arrNum {
		total += num
	}

	return total / float64(len(arrNum))
}
