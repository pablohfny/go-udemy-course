package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0664)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFileName)

	if err != nil {
		return 0, errors.New("failed to find balance file")
	}

	balanceText := string(data)

	fileBalance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("failed to parse balance file value")
	}

	return fileBalance, nil
}

var balance float64 = 0

func main() {
	var err error
	balance, err = getBalanceFromFile()

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

func displayMenu() {
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func obtainInput() int {
	var input int
	fmt.Print("Your input: ")
	fmt.Scan(&input)
	return input
}

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
	writeBalanceToFile(balance)
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
	writeBalanceToFile(balance)
}
