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
		buffer: make([]float64, n-1),
	}
}

func dot(n int, x []float64, y []float64) float64 {
	var result float64
	for i := 0; i < n; i++ {
		result += x[i] * y[i]
	}
	return result
}

func directI(fir *FirFilter, samples []float64) []float64 {
	length := fir.length
	coeffs := fir.coeffs
	output := make([]float64, len(samples))
	buffer := fir.buffer
	for i, x := range samples {
		// Shift the values
		result := x * coeffs[0]
		for j := length - 1; j >= 2; j-- {
			result += buffer[j-1] * coeffs[j]
			buffer[j-1] = buffer[j-2]
		}
		result += buffer[0] * coeffs[1]
		buffer[0] = x

		output[i] = result
	}

	return output
}

func directII(fir *FirFilter, samples []float64) []float64 {
	length := fir.length
	coeffs := fir.coeffs
	output := make([]float64, len(samples))
	buffer := fir.buffer
	for i, x := range samples {
		// Shift the values
		result := 0.0
		for j := length - 1; j > 0; j-- {
			result += buffer[j] * coeffs[j]
			buffer[j] = buffer[j-1]
		}
		result += buffer[0] * coeffs[0]
		buffer[0] = x

		output[i] = result
	}

	return output
}
