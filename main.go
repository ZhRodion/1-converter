package main

import "fmt"

const (
	usdToRub = 77.0
	usdToEur = 0.85
	eurToRub = usdToRub / usdToEur
)

func main() {
	for {
		initialCurrency := getInitialCurrency()
		currencyValue := currencyValue()
		targetCurrency := getTargetCurrency(initialCurrency)
		calculateCurrency := calculateCurrency(initialCurrency, currencyValue, targetCurrency)
		currencyOutput(calculateCurrency, targetCurrency)

		fmt.Println("Хотите продолжить? (y/n)")
		var continueInput string
		fmt.Scanln(&continueInput)
		if continueInput == "y" || continueInput == "Y" {
			continue
		} else {
			break
		}
	}
}

func getInitialCurrency() int {
	var currency int

	for {
		fmt.Println("Выберете вашу валюту: ")
		fmt.Println("1. Рубли")
		fmt.Println("2. Евро")
		fmt.Println("3. Доллары")
		fmt.Scanln(&currency)

		if currency > 0 && currency < 4 {
			return currency
		}
		fmt.Println("Введите корректный номер валюты")
	}
}

func currencyValue() float64 {
	var currencyDigit float64

	for {
		fmt.Println("Введите сумму валюты: ")
		fmt.Scanln(&currencyDigit)

		if currencyDigit > 0 {
			return currencyDigit
		}
		fmt.Println("Введите корректную сумму больше нуля")
	}
}

func getTargetCurrency(initialCurrency int) int {
	var targetCurrency int

	for {
		fmt.Println("Выберете валюту, в которую хотите перевести: ")

		switch initialCurrency {
		case 1:
			fmt.Println("1. Евро")
			fmt.Println("2. Доллары")
		case 2:
			fmt.Println("1. Рубли")
			fmt.Println("2. Доллары")
		case 3:
			fmt.Println("1. Рубли")
			fmt.Println("2. Евро")
		}

		fmt.Scanln(&targetCurrency)

		if targetCurrency >= 1 && targetCurrency <= 2 {
			switch initialCurrency {
			case 1: // Из рублей
				if targetCurrency == 1 {
					return 2 // Евро
				}
				return 3 // Доллары
			case 2: // Из евро
				if targetCurrency == 1 {
					return 1 // Рубли
				}
				return 3 // Доллары
			case 3: // Из долларов
				if targetCurrency == 1 {
					return 1 // Рубли
				}
				return 2 // Евро
			}
		}
		fmt.Println("Выберите корректный номер валюты (1-2)")
	}
}

func calculateCurrency(initialCurrency int, currencyValue float64, targetCurrency int) float64 {
	switch initialCurrency {
	case 1: // Из рублей
		switch targetCurrency {
		case 2: // в евро
			return currencyValue / eurToRub
		case 3: // в доллары
			return currencyValue / usdToRub
		}
	case 2: // Из евро
		switch targetCurrency {
		case 1: // в рубли
			return currencyValue * eurToRub
		case 3: // в доллары
			return currencyValue / usdToEur
		}
	case 3: // Из долларов
		switch targetCurrency {
		case 1: // в рубли
			return currencyValue * usdToRub
		case 2: // в евро
			return currencyValue * usdToEur
		}
	}
	return 0
}

func currencyOutput(calculateCurrency float64, targetCurrency int) {

	switch targetCurrency {
	case 1:
		fmt.Printf("Результат: %.2f рублей\n", calculateCurrency)
	case 2:
		fmt.Printf("Результат: %.2f евро\n", calculateCurrency)
	case 3:
		fmt.Printf("Результат: %.2f долларов\n", calculateCurrency)
	}
}
