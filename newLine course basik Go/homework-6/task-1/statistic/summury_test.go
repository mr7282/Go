package statistic

import (
	"testing"
)

func TestSum(t *testing.T) {
	var numSum float64
	numSum = Sum([]float64{1,2,5})
	if numSum != 9 {
		t.Error("expected 9, got", numSum)
	}
}