package main

import "fmt"

func main() {
	usdToEuro := 0.85
	usdToRub := 77.0

	euroToRub := usdToRub / usdToEuro

	fmt.Print(euroToRub)
}
