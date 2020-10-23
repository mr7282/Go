package statistic

import (
	"testing"
)

func TestAverage(t *testing.T) {
	var testNum float64
	testNum = Average([]float64{1, 2})
	if testNum != 1 {
		t.Error("expected 1, got", testNum)
	}
}
