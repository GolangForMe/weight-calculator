package main

import (
	"fmt"

	"github.com/gearmobile/calculator/calc"
	"github.com/gearmobile/calculator/check"
)

func main() {
	weight := 70
	height := 1.7
	imt := calc.Calc(weight, height)

	fmt.Printf("Индекс веса равен %f\n", imt)
	fmt.Printf(check.Check(imt))
}
