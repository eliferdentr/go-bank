package main

import (
	"fmt"
	"time"

	"elif.com/bank/fileoperations"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {

	var input float64
	var _, error = fileoperations.GetFloatFromFile(balanceFile)
	if error != nil {
		fmt.Printf("Error: %s\n", error)
		panic("A balance.txt file must exist before the program starts")
	}
	fmt.Println("Our phone number is: ",randomdata.PhoneNumber())
	for {
		showOptions()
		fmt.Scan(&input)

		switch input {
		case 1:
			checkBalance()
		case 2:
			depositMoney()
		case 3:
			withdrawMoney()
		case 4:
			fmt.Println("Exiting the program. Have a nice day!")
			return
		default:
			fmt.Println("Invalid input. Please enter a number between 1 - 4.")

		}

	}
}

func checkBalance() {
	var balance, error = fileoperations.GetFloatFromFile(balanceFile)
	if error != nil {
		fmt.Printf("Error: %s\n", error)
		return
	}
	fmt.Printf("Current Balance Is: %.2f\n", balance)
}


func depositMoney() {
	fmt.Println("Enter the value you want to deposit: ")
	var deposit float64
	fmt.Scan(&deposit)

	if deposit <= 0 {
		fmt.Println("Invalid amount. Please enter a positive number.")
	} else {
		var balance, error = fileoperations.GetFloatFromFile(balanceFile)
		if error != nil {
			fmt.Printf("Error: %s\n", error)
			return
		}
		balance += deposit
		fileoperations.WriteToFile(balance, balanceFile)
		fmt.Printf("Deposit successful. New balance is: %.2f\n", balance)
		time.Sleep(5 * time.Second)

	}
}
func withdrawMoney() {
	fmt.Println("Enter the value you want to withdraw: ")
	var withdrawAmount float64
	var balance, error = fileoperations.GetFloatFromFile(balanceFile)
	if error != nil {
		fmt.Printf("Error: %s\n", error)
		return
	}
	fmt.Scan(&withdrawAmount)
	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount. Please enter a positive number.")
	} else if withdrawAmount > balance {
		fmt.Println("Insufficient balance. You can not withdraw more than your current balance.")
	} else {
		balance -= withdrawAmount
		fileoperations.WriteToFile(balance, balanceFile)
		fmt.Printf("Withdrawal successful! New balance: %.2f\n", balance)
		time.Sleep(2 * time.Second)
	}
}

