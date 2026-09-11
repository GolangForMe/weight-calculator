package calc

import (
	"errors"
	"math"
)

// Calc - функция для расчета индекса массы тела
func Calc(weight int, height float64) (float64, error) {
	const base float64 = 2.0

	if weight <= 0 || height <= 0 {
		return 0, errors.New("Ошибка! Не указан вес или рост")
	}

	return float64(weight) / math.Pow(height/100, base), nil
}
