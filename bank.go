package main

import "fmt"
import "time"

var balance float64

func main() {
	
	formattedQuestion1 := fmt.Sprintf("What would you like to do? \n")
	formattedQuestion2 := fmt.Sprintf("1. Check balance. \n")
	formattedQuestion3 := fmt.Sprintf("2. Deposit money \n")
	formattedQuestion4 := fmt.Sprintf("3. Withdraw money \n")
	formattedQuestion5 := fmt.Sprintf("4. Exit \n")
	balance = 500

    var input
	for {
		fmt.Print(formattedQuestion1, formattedQuestion2, formattedQuestion3, formattedQuestion4, formattedQuestion5)
		fmt.Scan(&input)
		validInputCheck := checkInputValidity(input)
		if validInputCheck < -1 {
			fmt.Print("Input value is invalid")
			continue
		}

		if input == 1 {
			checkBalance()
		} else if input == 2 {
			depositMoney()

		} else if input == 3 {

		} else if input == 4 {
			break
		} else {
			fmt.Print("Input value is invalid")
			continue
		}


		

	}
}

func checkInputValidity(input) int {
	if input < 0 {
		return -1
	} else {
		return 1
	}
}

func checkBalance () {
	fmt.Printf("Balance: %v", balance)
}
func depositMoney() {
	fmt.Println("Enter the value you want to deposit: ")
	var deposit
	fmt.Scan(&deposit)
	validInputCheck := checkInputValidity(deposit)
	if validInputCheck < -1 {
		fmt.Println("Input value is invalid")
	} else {
		balance += deposit
		fmt.Printf("Current balance is: %")
		time.Sleep(8 * time.Second)

	}
}
