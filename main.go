package main

import (
	"fmt"
)

var currencyRates = map[string]float64{
	"rub": 1.0,
	"usd": 77.0,
	"eur": 90.0,
}

func main() {
	for {
		fmt.Println("=== Конвертер валют ===")

		from := getCurrency("Исходная валюта (rub/usd/eur): ")
		to := getCurrency("Целевая валюта (rub/usd/eur): ")
		amount := getAmount()

		result := convert(from, to, &amount)
		fmt.Printf("Результат: %.2f %s\n", result, to)

		fmt.Println("Хотите продолжить? (y/n)")
		var cont string
		fmt.Scanln(&cont)
		if cont != "y" && cont != "Y" {
			break
		}
	}
}

func getCurrency(prompt string) string {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)

		if _, exists := currencyRates[input]; exists {
			return input
		}
		fmt.Println("Неверный ввод. Попробуйте: rub, usd, eur")
	}
}

func getAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму: ")
		fmt.Scanln(&amount)
		if amount > 0 {
			return amount
		}
		fmt.Println("Введите сумму больше 0")
	}
}

func convert(from, to string, amount *float64) float64 {
	fromRate := currencyRates[from]
	toRate := currencyRates[to]

	return *amount * (fromRate / toRate)
}
