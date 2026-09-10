package helpers

import (
	"fmt"

	"github.com/gearmobile/calculator/check"
)

func GetUserInput() (int, float64) {
	var (
		weight int     = 80
		height float64 = 1.7
	)

	fmt.Print("Введите свой вес в килограммах (80): ")
	fmt.Scan(&weight)

	fmt.Print("Введите свой рост в сантиметрах (170): ")
	fmt.Scan(&height)

	return weight, height
}

func PrintImt(imt float64) {
	fmt.Printf("Индекс веса равен %0.2f\n", imt)
	fmt.Println(check.Check(imt))
}
