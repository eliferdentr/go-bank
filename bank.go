package main

import "fmt"
import "time"

var balance float64

func main() {

	balance = 500
	var input float64
	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check balance.")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
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
	fmt.Printf("Current Balance Is: %.2f", balance)
}

func depositMoney() {
	fmt.Println("Enter the value you want to deposit: ")
	var deposit float64
	fmt.Scan(&deposit)

	if deposit <= 0 {
		fmt.Println("Invalid amount. Please enter a positive number.")
	} else {
		balance += deposit
		fmt.Printf("Deposit successful. New balance is: %.2f\n", balance)
		time.Sleep(5 * time.Second)

	}
}
func withdrawMoney() {
	fmt.Println("Enter the value you want to withdraw: ")
	var withdrawAmount float64
	fmt.Scan(&withdrawAmount)
	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount. Please enter a positive number.")
	} else if withdrawAmount > balance {
		fmt.Println("Insufficient balance. You can not withdraw more than your current balance.")
	} else {
		balance -= withdrawAmount
		fmt.Printf("Withdrawal successful! New balance: %.2f\n", balance)
		time.Sleep(2 * time.Second)
	}
}
