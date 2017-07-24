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

func directI(fir *FirFilter, samples []float64) []float64 {
	length := fir.length
	coeffs := fir.coeffs
	output := make([]float64, len(samples))
	buffer := fir.buffer
	for i, x := range samples {
		// Shift and accumulate
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
		output[i] = buffer[length-2] + (x * coeffs[length-1])
		for j := length - 2; j >= 1; j-- {
			buffer[j] = buffer[j-1] + x*coeffs[j]
		}
		buffer[0] = x * coeffs[0]

	}

	return output
}
