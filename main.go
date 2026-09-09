package main

import (
	"fmt"

	"github.com/gearmobile/calculator/calc"
	"github.com/gearmobile/calculator/check"
)

func main() {
	var (
		weight int
		height float64
	)

	fmt.Print("Введите свой вес в килограммах (180): ")
	fmt.Scan(&weight)

	fmt.Print("Введите свой рост в метрах (1.7): ")
	fmt.Scan(&height)

	imt := calc.Calc(weight, height)

	fmt.Printf("Индекс веса равен %f\n", imt)
	fmt.Printf(check.Check(imt))
}
