package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, revenueErr := getInput("Revenue: ")
	expences, expenseErr := getInput("Expences: ")
	taxes, taxErr := getInput("Taxes: ")
	if revenueErr != nil || expenseErr != nil || taxErr != nil {
		fmt.Println(taxErr)
		return
	}

	ebt, profit, ratio := calculateProfit(revenue, expences, taxes)

	writeCalculationToFile(ebt, profit, ratio)

	printOutput("EBT: ", ebt)
	printOutput("Profit: ", profit)
	printOutput("Ratio: ", ratio)
}

func getInput(message string) (float64, error) {
	var input float64
	fmt.Print(message)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("invalid values entered")
	}

	return input, nil
}

func printOutput(message string, value float64) {
	fmt.Print(message)
	fmt.Printf("%.2f\n", value)
}

func calculateProfit(revenue, expences, taxes float64) (ebt, profit, ratio float64) {
	ebt = revenue - expences
	profit = ebt * (1 - taxes/100)
	ratio = ebt / profit

	return
}

func writeCalculationToFile(ebt, profit, ratio float64) {
	calculationStr := fmt.Sprintf("Earnings before tax: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile("profit_calculation.txt", []byte(calculationStr), 0644)
}
