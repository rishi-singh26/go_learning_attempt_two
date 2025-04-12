package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBanalnceFile = "balance.txt"

func main() {
	balance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		fmt.Println("----------------")
		// panic(err)
	}

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
		if depositAmount <= 0 {
			fmt.Println("Enter a valid ampount!")
		} else {
			balance += depositAmount
			fmt.Printf("Your updated balance is %.4f\n\n", balance)
			writeBalanceToFile(balance)
		}
		startBank(balance)
	case 3:
		fmt.Print("Enter the amount to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount <= 0 {
			fmt.Println("Enter a valid ampount!")
		} else if withdrawAmount > balance {
			fmt.Printf("Not enough balance, your current balance is %.4f\n\n", balance)
		} else {
			balance -= withdrawAmount
			fmt.Printf("Your updated balance is %.4f\n\n", balance)
			writeBalanceToFile(balance)
		}
		startBank(balance)
	default:
		fmt.Println("Bye bye")
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

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBanalnceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBanalnceFile)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance")
	}

	return balance, nil
}
