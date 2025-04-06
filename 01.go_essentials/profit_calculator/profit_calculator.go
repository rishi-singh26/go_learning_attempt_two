package main

import "fmt"

func main() {
	revenue := 1000.0
	expences := 500.0
	taxes := 7.0 // 25.6 => 0.256

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expences: ")
	fmt.Scan(&expences)

	fmt.Print("Taxes: ")
	fmt.Scan(&taxes)

	ebt := revenue - expences

	profit := ebt * (1 - taxes/100)

	ratio := ebt / profit

	fmt.Print("EBT: ")
	fmt.Println(ebt)

	fmt.Print("Profit: ")
	fmt.Println(profit)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}
