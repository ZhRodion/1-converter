package main

import "fmt"

func main() {
	usdToEuro := 0.85
	usdToRub := 77.0

	currencyInput := currencyInput()
	convertedValue := currencyConverter(currencyInput, usdToEuro, usdToRub)
	fmt.Printf("%.2f", convertedValue)
}

func currencyInput() int {
	var currency int

	fmt.Println("Введите сумму в рублях: ")
	fmt.Scanln(&currency)

	if currency > 0 {
		return currency
	} else {
		fmt.Println("Введите корректную сумму")

		return currencyInput()
	}
}

func currencyConverter(currencyValue int, currentCurrеncy, outputCurrency float64) float64 {
	return 1.22
}
