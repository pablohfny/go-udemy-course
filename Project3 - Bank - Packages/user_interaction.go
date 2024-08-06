package main

import "fmt"

func displayMenu() {
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func processInput(input int) {
	switch input {
	case 1:
		checkBalance()
	case 2:
		deposit()
	case 3:
		withdraw()
	case 4:
		fmt.Println("Good bye!")
	default:
		fmt.Println("Invalid input!")
	}
}

func obtainInput() int {
	var input int
	fmt.Print("Your input: ")
	fmt.Scan(&input)
	return input
}
