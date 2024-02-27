package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmmount float64
	var years float64
	var expectedReturnRate float64

	//Ask for user input
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmmount)
	fmt.Print("Investment years: ")
	fmt.Scan(&years)
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	//Variables declared with function
	futureValue, futureRealValue := calculateFutureValues(investmentAmmount, expectedReturnRate, years)

	// calculations done at variable declaration
	// futureValue := investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Standard Print Lines
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)

	//Formated Print Line
	// fmt.Printf("Future Value: $%.2f\nFuture Value (adjusted for Inflation): $%.2f", futureValue, futureRealValue)

	// Formated Print "Real" line
	// fmt.Printf(`Future Value: $%.2f
	// Future Value (adjusted for Inflation): $%.2f`, futureValue, futureRealValue)

	//Variable formated strings
	formatedFV := fmt.Sprintf("Future Value: $%.2f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): $%.2f\n", futureRealValue)
	fmt.Print(formatedFV, formatedRFV)
}

func calculateFutureValues(investmentAmmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}
