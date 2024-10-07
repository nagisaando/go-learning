package main

import (
	"fmt"

	"example.com/bank/fileops"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	balance, _ := fileops.ReadFloatFromFile(accountBalanceFile, 1000)

	// if err != nil {
	// 	panic(err) // this will stop the further execution.
	// }

	fmt.Println("Welcome!")
	fmt.Printf("current balance is %v \n", balance)
	fmt.Println("Reach us at", randomdata.PhoneNumber())

	for {
		var choice int

		// this function is coming from communication.go.
		// we can use this file without importing or exporting because communication.go is part of main package.
		printOption()
		_, err := fmt.Scanln(&choice)

		// nil means absence of useful value
		if err != nil {
			fmt.Println("Invalid value, please try again")
			var discard string

			// this is to clear buffer, because fmt.Scanln will leave the invalid value in buffer, and waiting to be processed. It has to be cleared by reading and storing the invalid value
			fmt.Scanln(&discard)
			continue // skips further execution but keep the loop
		}

		if choice == 4 {
			fmt.Println("Goodbye!")
			break // to stops the loop
		}
		balance = modifyCurrentBalance(choice, balance)
		fileops.WriteFloatFromFile(accountBalanceFile, balance)

	}

}

func modifyCurrentBalance(choice int, balance float64) float64 {
	var currentBalance = balance

	switch choice {
	case 1:
		fmt.Printf("Your balance is $%v\n", currentBalance)
	case 2:
		var deposit float64
		fmt.Println("How much do you want to deposit?")
		fmt.Scanln(&deposit)

		currentBalance += deposit
		fmt.Printf("Your balance is $%v\n", currentBalance)

	case 3:
		var withdrawAmount float64
		fmt.Println("How much do u want to withdraw?")
		_, err := fmt.Scanln(&withdrawAmount)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			var discard string
			fmt.Scanln(&discard)
			return currentBalance // return stops further execution
		}
		if withdrawAmount < 0 {
			fmt.Printf("Please add positive number")
		} else if currentBalance < withdrawAmount {
			fmt.Println("Current balance is lower than withdraw amount. Please adjust the amount to withdraw and try again")
		} else {
			currentBalance -= withdrawAmount
			fmt.Printf("Your balance is $%v\n", currentBalance)

		}

	default:
		fmt.Println("Invalid option, please try again.")

	}

	return currentBalance
}
