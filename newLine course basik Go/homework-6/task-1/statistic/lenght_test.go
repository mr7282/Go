package statistic

import "testing"

func TestLen(t *testing.T) {
	var NumLen int
	NumLen = Len([]float64{1, 1, 1})
	if NumLen != 4 {
		t.Error("Expected 4, got", NumLen)
	}
}
