package main

import "fmt"

func main() {
	revenue := getInput("Revenue: ")
	expences := getInput("Expences: ")
	taxes := getInput("Taxes: ")

	ebt, profit, ratio := calculateProfit(revenue, expences, taxes)

	printOutput("EBT: ", ebt)
	printOutput("Profit: ", profit)
	printOutput("Ratio: ", ratio)
}

func getInput(message string) (input float64) {
	fmt.Print(message)
	fmt.Scan(&input)
	return
}

func printOutput(message string, value float64) {
	fmt.Print(message)
	fmt.Println(value)
}

func calculateProfit(revenue, expences, taxes float64) (ebt, profit, ratio float64) {
	ebt = revenue - expences
	profit = ebt * (1 - taxes/100)
	ratio = ebt / profit

	return
}
