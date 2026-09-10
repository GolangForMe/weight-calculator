package calc

import "math"

// Calc - функция для расчета индекса массы тела
func Calc(weight int, height float64) float64 {
	const base float64 = 2.0
	return float64(weight) / math.Pow(height/100, base)
}
