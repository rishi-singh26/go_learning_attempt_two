package main

import "fmt"

func main() {
	balance := 1000.0
	fmt.Println("Welcome to the Bank!")
	startBank(balance)
}

func startBank(balance float64) {
	input := getInput()
	switch input {
	case 1:
		fmt.Printf("Your balance is %.4f\n\n", balance)
		startBank(balance)
	case 2:
		fmt.Print("Enter deposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		balance += depositAmount
		fmt.Printf("Your updated balance is %.4f\n\n", balance)
		startBank(balance)
	case 3:
		fmt.Print("Enter the amount to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > balance {
			fmt.Printf("Not enough balance, your current balance is %.4f\n\n", balance)
		} else {
			balance -= withdrawAmount
			fmt.Printf("Your updated balance is %.4f\n\n", balance)
		}
		startBank(balance)
	default:
		fmt.Print("Bye bye\n")
	}
}

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

// func depositMoney(amount float64, balance *float64) {
// 	balance
// }
