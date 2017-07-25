package fir

import (
	"fmt"
	"testing"
)

func TestFirFilter(t *testing.T) {
	fir := NewFirFilter([]float64{0, 0, 1, 0, 0})
	out := directII(fir, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(out)
	out = directII(fir, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(out)
}

func BenchmarkI(b *testing.B) {
	fir := NewFirFilter(make([]float64, 1000))
	for i := 0; i < b.N; i++ {
		directI(fir, make([]float64, 1))
	}
}

func BenchmarkII(b *testing.B) {
	fir := NewFirFilter(make([]float64, 1000))

	for i := 0; i < b.N; i++ {
		directII(fir, make([]float64, 1))
	}
}
