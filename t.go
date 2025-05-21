package main

import (
	"fmt"
	"strings"
)

const (
	usdToEur = 0.94
	usdToRub = 75.5
	eurToUsd = 1 / usdToEur
	rubToUsd = 1 / usdToRub
)

func main() {
	fmt.Println("=== Конвертер валют ===")

	// Шаг 1: Ввод исходной валюты
	fromCurrency := readCurrency("Введите исходную валюту (USD, EUR, RUB): ")

	// Шаг 2: Ввод суммы
	amount := readNumber(fmt.Sprintf("Введите сумму в %s для конвертации: ", fromCurrency))

	// Шаг 3: Ввод целевой валюты
	toCurrency := readCurrency("Введите целевую валюту (USD, EUR, RUB): ")

	// Проверка, что исходная и целевая валюты не совпадают
	for strings.EqualFold(fromCurrency, toCurrency) {
		fmt.Println("Исходная и целевая валюта не должны совпадать. Попробуйте снова.")
		toCurrency = readCurrency("Введите целевую валюту (USD, EUR, RUB): ")
	}

	// Шаг 4: Вычисление и вывод результата
	result := convertCurrency(amount, fromCurrency, toCurrency)
	fmt.Printf("%.2f %s = %.2f %s\n", amount, strings.ToUpper(fromCurrency), result, strings.ToUpper(toCurrency))
}

// readCurrency запрашивает у пользователя валюту и проверяет корректность
func readCurrency(prompt string) string {
	validCurrencies := []string{"USD", "EUR", "RUB"}
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scan(&input)
		input = strings.ToUpper(input)

		for _, v := range validCurrencies {
			if input == v {
				return input
			}
		}
		fmt.Println("Неверная валюта. Пожалуйста, выберите из: USD, EUR, RUB.")
	}
}

// readNumber запрашивает у пользователя число с проверкой корректности
func readNumber(prompt string) float64 {
	var value float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&value)
		if err != nil || value < 0 {
			fmt.Println("Пожалуйста, введите корректное положительное число.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return value
	}
}

// convertCurrency конвертирует сумму из одной валюты в другую
func convertCurrency(amount float64, from, to string) float64 {
	switch from {
	case "USD":
		switch to {
		case "EUR":
			return amount * usdToEur
		case "RUB":
			return amount * usdToRub
		}
	case "EUR":
		switch to {
		case "USD":
			return amount * eurToUsd
		case "RUB":
			// EUR -> USD -> RUB
			usd := amount * eurToUsd
			return usd * usdToRub
		}
	case "RUB":
		switch to {
		case "USD":
			return amount * rubToUsd
		case "EUR":
			// RUB -> USD -> EUR
			usd := amount * rubToUsd
			return usd * usdToEur
		}
	}
	// Если что-то пошло не так, возвращаем 0
	return 0
}
