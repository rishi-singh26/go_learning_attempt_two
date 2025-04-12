package main

import "fmt"

func getInput() (input float64) {
	fmt.Print(`What do you wnat to do?
1. Check Balance
2. Deposit Money
3. Withdraw Money
4. Exit
Enter Your Input: `)
	fmt.Scan(&input)
	return
}
