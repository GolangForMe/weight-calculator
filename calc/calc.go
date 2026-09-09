package calc

import "math"

func Calc(weight int, height float64) float64 {
	base := 2.0
	return float64(weight) / math.Pow(height, base)
}
