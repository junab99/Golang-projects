package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)


const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Amount must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Print("Your withdrawl amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Amount must be greater than zero.")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Insufficent funds")
				continue
			}

			accountBalance -= withdrawlAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}

		// wantsCheckBalance := choice == 1

		// if choice == 1 {
		// 	fmt.Println("Your balance is ", accountBalance)
		// } else if choice == 2{
		// 	fmt.Print("Your deposit amount: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Amount must be greater than zero.")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else if choice == 3{
		// 	fmt.Print("Your withdrawl amount: ")
		// 	var withdrawlAmount float64
		// 	fmt.Scan(&withdrawlAmount)

		// 	if withdrawlAmount <= 0 {
		// 		fmt.Println("Amount must be greater than zero.")
		// 		continue
		// 	}

		// 	if withdrawlAmount > accountBalance {
		// 		fmt.Println("Insufficent funds")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawlAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }

	}

}
