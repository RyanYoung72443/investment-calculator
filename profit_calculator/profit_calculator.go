package main

import (
	"fmt"
)

func main() {
	revenue := getValueFromUser("Earned revenue: ")
	expenses := getValueFromUser("Expenses: ")
	taxRate := getValueFromUser("Tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	formatedEBT := fmt.Sprintf("EBT: $%.2f\n", ebt)
	formatedProfit := fmt.Sprintf("Profit: $%.2f\n", profit)
	formatedRatio := fmt.Sprintf("Ratio: %.2f%%\n", ratio)

	fmt.Print(formatedEBT, formatedProfit, formatedRatio)
}

func getValueFromUser(valueDesired string) (res float64) {
	fmt.Print(valueDesired)
	fmt.Scan(&res)
	return
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
