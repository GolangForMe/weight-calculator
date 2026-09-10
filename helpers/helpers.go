package helpers

import (
	"fmt"

	"github.com/gearmobile/calculator/check"
)

// GetUserInput запрашивает у пользователя свой вес и рост, а затем возвращает их в виде целого числа и float64
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

// PrintImt выводит индекс массы тела (IMT) и результат проверки на соответствие стандартам
func PrintImt(imt float64) {
	fmt.Println("---------------------")
	fmt.Printf("Индекс веса равен %0.2f\n", imt)
	fmt.Println(check.Check(imt))
}

// Проверка выбора пользователя
func CheckUserChoise() bool {
	var choise string
	fmt.Println("=======================")
	fmt.Print("Хотите продолжить (yes/no): ")
	fmt.Scan(&choise)

	if choise == "no" || choise == "n" {
		return true
	}

	return false
}
