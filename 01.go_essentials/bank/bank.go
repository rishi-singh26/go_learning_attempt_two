package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"rishisingh.in/bank/fileops"
)

const accountBanalnceFile = "balance.txt"

func main() {
	balance, err := fileops.ReadFloatFromFile(accountBanalnceFile)

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		fmt.Println("----------------")
		// panic(err)
	}

	fmt.Println("Welcome to the Bank!")
	fmt.Println(randomdata.PhoneNumber())
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
			fileops.WriteFloatToFile(balance, accountBanalnceFile)
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
			fileops.WriteFloatToFile(balance, accountBanalnceFile)
		}
		startBank(balance)
	default:
		fmt.Println("Bye bye")
	}
}
