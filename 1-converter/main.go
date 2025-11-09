package main

import "fmt"

func main() {
	const (
		USD_TO_EUR = 0.85
		USD_TO_RUB = 75.50
	)

	const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

	fmt.Printf("Курсы конвертации:\n")
	fmt.Printf("USD to EUR: %.4f\n", USD_TO_EUR)
	fmt.Printf("USD to RUB: %.4f\n", USD_TO_RUB)
	fmt.Printf("EUR to RUB: %.4f\n", EUR_TO_RUB)

	usdAmount := 100.0
	eurAmount := usdAmount * USD_TO_EUR
	rubAmount := usdAmount * USD_TO_RUB

	fmt.Printf("\nПример конвертации $%.2f:\n", usdAmount)
	fmt.Printf("В EUR: €%.2f\n", eurAmount)
	fmt.Printf("В RUB: ₽%.2f\n", rubAmount)

	eurAmount2 := 50.0
	rubAmount2 := eurAmount2 * EUR_TO_RUB

	fmt.Printf("\nПример конвертации €%.2f:\n", eurAmount2)
	fmt.Printf("В RUB: ₽%.2f\n", rubAmount2)
}
