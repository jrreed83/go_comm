package fir

type FirFilter struct {
	length int
	coeffs []float64
	buffer []float64
}

func NewFirFilter(coeffs []float64) *FirFilter {
	n := len(coeffs)
	c := make([]float64, n)
	copy(c, coeffs)
	return &FirFilter{
		length: n,
		coeffs: c,
		buffer: make([]float64, n),
	}
}

func dot(n int, x []float64, y []float64) float64 {
	var result float64
	for i := 0; i < n; i++ {
		result += x[i] * y[i]
	}
	return result
}

func step(fir *FirFilter, samples []float64) []float64 {
	length := fir.length
	coeffs := fir.coeffs
	output := make([]float64, len(samples))
	buffer := fir.buffer
	for i, x := range samples {
		output[i] = dot(length, coeffs, buffer)

		// Shift the values
		for j := length - 1; j > 0; j-- {
			buffer[j] = buffer[j-1]
		}
		buffer[0] = x
	}

	return output
}
