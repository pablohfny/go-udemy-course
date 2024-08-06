package main

import (
	"bank/fileops"
	"fmt"
)

const balanceFileName = "balance.txt"

var balance float64 = 0

func main() {
	var err error
	balance, err = fileops.GetFloatFromFile(balanceFileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	for input := 0; input != 4; {
		displayMenu()
		input = obtainInput()
		processInput(input)
	}
}

// func processInput(input int) {
// 	if input == 1 {
// 		checkBalance()
// 	} else if input == 2 {
// 		deposit()
// 	} else if input == 3 {
// 		withdraw()
// 	} else if input == 4 {
// 		fmt.Println("Good bye!")
// 	} else {
// 		fmt.Println("Invalid input!")
// 	}
// }

func checkBalance() {
	fmt.Println("Your Balance: ", balance)
}

func deposit() {
	var amount float64

	fmt.Print("Amount to deposit: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Amount should be greater than 0!")
		return
	}

	balance += amount
	fmt.Println("New Balance: ", balance)
	fileops.WriteFloatToFile(balance, balanceFileName)
}

func withdraw() {
	var amount float64

	fmt.Print("Amount to withdraw: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Amount should be greater than 0!")
		return
	}

	if amount > balance {
		fmt.Println("Insuficient funds, total balance: ", balance)
		return
	}

	balance -= amount
	fmt.Println("New Balance: ", balance)
	fileops.WriteFloatToFile(balance, balanceFileName)
}
