package fir

import (
	"fmt"
	"testing"
)

func TestFirFilter(t *testing.T) {
	fir := NewFirFilter([]float64{0, 1, 0})
	out := step(fir, []float64{0, 1, 1, 1, 2, 3, 1, 2, 3, 5, 3})
	fmt.Println(out)

}
