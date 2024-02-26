package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmenAmmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmenAmmount)
	fmt.Print("Investment years: ")
	fmt.Scan(&years)
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmenAmmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
