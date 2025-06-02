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
		currencyOutput(calculateCurrency)

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

	fmt.Println("Выберете вашу валюту: ")
	fmt.Println("1. Рубли")
	fmt.Println("2. Евро")
	fmt.Println("3. Доллары")
	fmt.Scanln(&currency)

	if currency > 0 && currency < 4 {
		return currency
	} else {
		fmt.Println("Введите корректный номер валюты")
		return getInitialCurrency()
	}
}

func currencyValue() int {
	var currencyDigit int

	fmt.Println("Введите сумму валюты: ")
	fmt.Scanln(&currencyDigit)

	switch {
	case currencyDigit <= 0:
		{
			fmt.Println("Введите корректную сумму больше нуля")
			return currencyValue()
		}

	case currencyDigit > 0:
		{
			return currencyDigit
		}
	}

	return currencyDigit
}

func getTargetCurrency(initialCurrency int) int {
	var targetCurrency int

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

	if targetCurrency < 1 || targetCurrency > 2 {
		fmt.Println("Выберите корректный номер валюты (1-2)")
		return getTargetCurrency(initialCurrency)
	}

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

	return targetCurrency
}

func calculateCurrency(initialCurrency int, currencyValue int, targetCurrency int) float64 {
	switch initialCurrency {
	case 1: // Из рублей
		switch targetCurrency {
		case 2: // в евро
			return float64(currencyValue) / eurToRub
		case 3: // в доллары
			return float64(currencyValue) / usdToRub
		}
	case 2: // Из евро
		switch targetCurrency {
		case 1: // в рубли
			return float64(currencyValue) * eurToRub
		case 3: // в доллары
			return float64(currencyValue) / usdToEur
		}
	case 3: // Из долларов
		switch targetCurrency {
		case 1: // в рубли
			return float64(currencyValue) * usdToRub
		case 2: // в евро
			return float64(currencyValue) * usdToEur
		}
	}
	return 0
}

func currencyOutput(calculateCurrency float64) {
	fmt.Printf("Результат: %.2f\n", calculateCurrency)
}
