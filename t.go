package main

import (
	"fmt"
)

func main() {
	const usdToEur = 0.94
	const usdToRub = 75.5
	eurToRub := usdToRub / usdToEur
	fmt.Println(eurToRub)
	// step 1
	userInput()
}
func userInput() {
	var input float64
	fmt.Println("Введите сумму, которую вы хотели бы конвертировать: ")
	fmt.Scan(&input)
}
func calculation() {

}
