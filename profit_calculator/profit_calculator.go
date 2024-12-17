package main

import (
	"fmt"
	"errors"
	"os"
)

//Goals
// 1) Validate user input
// 		=> Show error message & exit if invalid input is provided
// 		- No Negative numbers
// 		- Not 0
// 	2) Store calculated results into file
// func main() {

func main() {

	revenue, err1:= getUserInput("Revenue: ")

	// if err != nil {
	// 	panic(err)
	// }

	expenses, err2 := getUserInput("Expenses: ")

	// if err != nil {
	// 	panic(err)
	// }


	taxRate, err3 := getUserInput("Tax Rate: ")

	// if err != nil {
	// 	panic(err)
	// }


	if err1 != nil || err2 != nil || err3 !=nil {
		panic(err1)
	}

	ebt, profit, ratio := calculations(revenue, expenses, taxRate)

	storeResults(ebt, profit, ratio)

	fmt.Printf("Earnings before tax: %v\n", ebt)

	fmt.Println("Profit:", profit)

	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(formattedRatio)

}

func storeResults(ebt, profit, ratio float64) {

	results := fmt.Sprintf("EBT: %.2f\nPROFIT: %.1f\nRATIO: %.2f\n", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)

}


func getUserInput(infoText string) (float64, error) {

	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("input must be greater than 0")
	}

	return userInput, nil
}


func calculations(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {

	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
