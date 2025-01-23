package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main (){
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------------------------")
		panic("sorry, cannot continue")
	}

	fmt.Println("Welcome to GO Bank")

	for {
		fmt.Println("What do you do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposite Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1: 
			fmt.Println("Your Balance: ", accountBalance)
		case 2: 
			var depositMoney float64
			fmt.Scan(&depositMoney)
			if depositMoney <= 0 {
				fmt.Println("Deposit Money must be greater than 0")
				continue
			}
			accountBalance += depositMoney
			fmt.Println("Balance Updated! New Amount:", accountBalance)
			fileops.WriteValueToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawMoney float64
			fmt.Scan(&withdrawMoney)
			if withdrawMoney <= 0 {
				fmt.Println("Withdraw Money must be greater than 0")
				continue
			}

			if withdrawMoney > accountBalance {
				fmt.Println("Withdraw Money must be Less than or Equal to Account Balance")
				continue
			}

			accountBalance -= withdrawMoney
			fmt.Println("Balance Updated! New Amount:", accountBalance)
			fileops.WriteValueToFile(accountBalance, accountBalanceFile)
		case 4:
			return
		default:
			fmt.Println("You must choose a number from 1 to 4")
			continue
		}
	}
}