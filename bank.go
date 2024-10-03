package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)


const balanceFile = "balance.txt"

func main() {

	var input float64
	var _, error = getBalance()
	if error != nil {
		fmt.Printf("Error: %s\n", error)
		panic("A balance.txt file must exist before the program starts")
	}
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
func checkBalance () {
	var balance, error = getBalance()
	if error != nil {
		fmt.Printf("Error: %s\n", error)
		return
	}
	fmt.Printf("Current Balance Is: %.2f\n", balance)
}


func getBalance() (float64, error){
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000, errors.New("Failed to read file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}
	return balance, nil

}
func depositMoney() {
	fmt.Println("Enter the value you want to deposit: ")
	var deposit float64
	fmt.Scan(&deposit)

	if deposit <= 0 {
		fmt.Println("Invalid amount. Please enter a positive number.")
	} else {
		var balance, error = getBalance()
		if error != nil {
			fmt.Printf("Error: %s\n", error)
			return
		}
		balance += deposit
		writeBalanceToFile(balance)
		fmt.Printf("Deposit successful. New balance is: %.2f\n", balance)
		time.Sleep(5 * time.Second)

	}
}
func withdrawMoney() {
	fmt.Println("Enter the value you want to withdraw: ")
	var withdrawAmount float64
	var balance, error = getBalance()
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
		writeBalanceToFile(balance)
		fmt.Printf("Withdrawal successful! New balance: %.2f\n", balance)
		time.Sleep(2 * time.Second)
	}
}
func writeBalanceToFile(balance float64) {
	// Balance'ı string'e çevirirken format belirliyoruz
	balanceText := fmt.Sprintf("%.2f", balance)
	err := os.WriteFile(balanceFile, []byte(balanceText), 0644) // Hata kontrolü ekliyoruz
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Balance successfully written to file.")
}


