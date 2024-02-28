package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getValueFromUser("Earned revenue: ")
	if err != nil {
		fmt.Printf("%f %v", revenue, err)
		return
	}
	expenses, err := getValueFromUser("Expenses: ")
	if err != nil {
		fmt.Printf("%f %v", expenses, err)
		return
	}
	taxRate, err := getValueFromUser("Tax rate: ")
	if err != nil {
		fmt.Printf("%f %v", taxRate, err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	results := writeCalculationToFile(ebt, profit, ratio)

	fmt.Print(results)
}

func writeCalculationToFile(ebt, profit, ratio float64) string {
	document := fmt.Sprintf("EBT: $%.2f\nProfit: $%.2f\nRatio: %.2f%%\n", ebt, profit, ratio)
	os.WriteFile("calculation.txt", []byte(document), 0644)
	return document
}

func getValueFromUser(valueDesired string) (res float64, err error) {
	fmt.Print(valueDesired)
	fmt.Scan(&res)
	if res <= 0 {
		return res, errors.New("value must be a positive number")
	}
	return res, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
